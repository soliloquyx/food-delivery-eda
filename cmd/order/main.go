package main

import (
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
	grpcin "github.com/soliloquyx/food-delivery-eda/internal/order/adapters/in/grpc"
	"github.com/soliloquyx/food-delivery-eda/internal/order/app"
	"google.golang.org/grpc"
)

func main() {
	svc := app.NewService()
	grpcServer := grpc.NewServer()

	orderv1.RegisterOrderServiceServer(grpcServer, grpcin.NewServer(svc))
}
