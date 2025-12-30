package grpcin

import (
	"context"

	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
	"github.com/soliloquyx/food-delivery-eda/internal/order/adapters/grpcin/placeorder"
	"github.com/soliloquyx/food-delivery-eda/internal/order/order"
)

type server struct {
	orderv1.UnimplementedOrderServiceServer
	svc Service
}

type Service interface {
	PlaceOrder(ctx context.Context, in order.PlaceOrderInput) (order.PlaceOrderResult, error)
}

func NewServer(s Service) orderv1.OrderServiceServer {
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
