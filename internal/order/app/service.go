package app

import (
	"context"

	"github.com/soliloquyx/food-delivery-eda/internal/order/ports"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) PlaceOrder(ctx context.Context, in ports.PlaceOrderInput) (ports.PlaceOrderResult, error) {
	return ports.PlaceOrderResult{
		OrderID: 1,
		Status:  ports.StatusConfirmed,
	}, nil
}
