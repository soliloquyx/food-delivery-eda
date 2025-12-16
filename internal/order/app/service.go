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
