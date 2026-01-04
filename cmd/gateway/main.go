package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/httpin"
	"github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/httpin/middleware"
	"github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/orderclient"
	"github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
	"github.com/soliloquyx/food-delivery-eda/internal/gateway/config"
	"github.com/soliloquyx/food-delivery-eda/internal/telemetry"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/zap"
)

func run(ctx context.Context) error {
	cfg, err := config.FromEnv()
	if err != nil {
		return err
	}

	orderClient, cleanup, err := orderclient.New(cfg.OrderGRPCAddr)
	if err != nil {
		return err
	}
	defer cleanup()

	shutdown, err := telemetry.InitOTLP(
		ctx,
		telemetry.Config{
			SvcName:  cfg.SvcName,
			Endpoint: cfg.OTLPEndpoint,
		},
	)
	if err != nil {
		return err
	}
	defer shutdown(ctx)

	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()

	orderSvc := order.NewService(orderClient)

	mw := middleware.Chain{
		otelhttp.NewMiddleware(cfg.SvcName),
		middleware.RequestID,
	}

	httpHandler := httpin.NewHandler(orderSvc)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /orders", httpin.Adapt(logger, httpHandler.PlaceOrder))

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           mw.Then(mux),
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("HTTP server listening", zap.String("addr", cfg.HTTPAddr))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		logger.Info("shutdown started")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			return err
		}

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
