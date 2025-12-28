package main

import (
	"context"
	"log"
	"net"
	"os/signal"
	"syscall"

	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
	grpcin "github.com/soliloquyx/food-delivery-eda/internal/order/adapters/in/grpc"
	"github.com/soliloquyx/food-delivery-eda/internal/order/adapters/out/postgres"
	"github.com/soliloquyx/food-delivery-eda/internal/order/config"
	"github.com/soliloquyx/food-delivery-eda/internal/order/order"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func run(ctx context.Context) error {
	cfg, err := config.FromEnv()
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		return err
	}

	db, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}

	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	orderRepo := postgres.NewOrderRepo(db)
	svc := order.NewService(&orderRepo)
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	orderv1.RegisterOrderServiceServer(grpcServer, grpcin.NewServer(svc))

	errCh := make(chan error, 1)
	go func() {
		logger.Info("gRPC server listening", zap.String("addr", cfg.GRPCAddr))
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		logger.Info("shutdown started")
		grpcServer.GracefulStop()
		logger.Info("shutdown complete")
	case err := <-errCh:
		return err
	}

	return nil
}

func main() {
	appCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := run(appCtx); err != nil {
		log.Fatal(err)
	}
}
