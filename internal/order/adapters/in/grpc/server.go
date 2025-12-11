package grpc

import (
	"context"

	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
	"github.com/soliloquyx/food-delivery-eda/internal/order/ports"
)

type server struct {
	orderv1.UnimplementedOrderServiceServer
	svc ports.Service
}

func NewServer(s ports.Service) orderv1.OrderServiceServer {
	return &server{
		svc: s,
	}
}

func (s *server) PlaceOrder(context.Context, *orderv1.PlaceOrderRequest) (*orderv1.PlaceOrderResponse, error)
