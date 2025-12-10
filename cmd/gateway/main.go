package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpin "github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/in/http"
	orderout "github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/out/order"
	orderapp "github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
)

func run(ctx context.Context) error {
	svcName := "gateway"

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	orderClient, cleanup, err := orderout.NewGRPC("")
	if err != nil {
		return err
	}
	defer cleanup()
	orderSvc := orderapp.NewService(orderClient)
	httpHandler := httpin.NewHandler(orderSvc)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /orders", httpHandler.PlaceOrder)

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		log.Printf("%s: listening on port %s", svcName, port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Printf("%s: shutdown signal received", svcName)

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}

		log.Printf("%s: graceful shutdown complete", svcName)
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
