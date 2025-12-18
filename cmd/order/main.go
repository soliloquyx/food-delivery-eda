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
	"github.com/soliloquyx/food-delivery-eda/internal/order/app"
	"github.com/soliloquyx/food-delivery-eda/internal/order/config"
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

	orderRepo := postgres.NewOrderRepo(db)
	svc := app.NewService(&orderRepo)
	grpcServer := grpc.NewServer()
	orderv1.RegisterOrderServiceServer(grpcServer, grpcin.NewServer(svc))

	errCh := make(chan error, 1)
	go func() {
		log.Printf("%s: gRPC listening on %s", cfg.SvcName, cfg.GRPCAddr)
		if err := grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Printf("%s: shutdown signal received", cfg.SvcName)
		grpcServer.GracefulStop()
		log.Printf("%s: graceful shutdown complete", cfg.SvcName)
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
