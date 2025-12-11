package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	httpin "github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/in/http"
	orderout "github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/out/order"
	orderapp "github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
	"github.com/soliloquyx/food-delivery-eda/internal/gateway/config"
)

func run(ctx context.Context) error {
	cfg, err := config.FromEnv()
	if err != nil {
		return err
	}

	orderClient, cleanup, err := orderout.NewClient(cfg.OrderGRPCAddr)
	if err != nil {
		return err
	}
	defer cleanup()

	orderSvc := orderapp.NewService(orderClient)
	httpHandler := httpin.NewHandler(orderSvc)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /orders", httpHandler.PlaceOrder)

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Printf("%s: listening on addr %s", cfg.SvcName, cfg.HTTPAddr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Printf("%s: shutdown signal received", cfg.SvcName)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}

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
