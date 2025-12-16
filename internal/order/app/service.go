package app

import (
	"context"
)

type Service interface {
	PlaceOrder(ctx context.Context, in PlaceOrderInput) (PlaceOrderResult, error)
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) PlaceOrder(ctx context.Context, in PlaceOrderInput) (PlaceOrderResult, error) {
	return PlaceOrderResult{
		OrderID: 1,
		Status:  StatusConfirmed,
	}, nil
}
