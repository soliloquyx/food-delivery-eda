package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/in/httpapi"
	order "github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
)

func run(ctx context.Context) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	orderSvc := order.NewService()
	httpHandler := httpapi.NewHandler(orderSvc)
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
		log.Printf("api gateway: listening on port %s", port)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		log.Println("api gateway: shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}

		log.Println("api gateway: graceful shutdown complete")
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
