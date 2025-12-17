package grpc

import (
	"context"

	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
	"github.com/soliloquyx/food-delivery-eda/internal/order/adapters/in/grpc/placeorder"
	"github.com/soliloquyx/food-delivery-eda/internal/order/app"
)

type server struct {
	orderv1.UnimplementedOrderServiceServer
	svc app.Service
}

func NewServer(s app.Service) orderv1.OrderServiceServer {
	return &server{
		svc: s,
	}
}

func (s *server) PlaceOrder(ctx context.Context, req *orderv1.PlaceOrderRequest) (*orderv1.PlaceOrderResponse, error) {
	in, err := placeorder.ToInput(req)
	if err != nil {
		return nil, err
	}

	result, err := s.svc.PlaceOrder(ctx, in)
	if err != nil {
		return nil, err
	}

	return placeorder.ToResponse(result), nil
}
