package order

import (
	"context"
)

type Service interface {
	PlaceOrder(ctx context.Context, in PlaceInput) (PlaceResult, error)
}

type Client interface {
	PlaceOrder(ctx context.Context, in PlaceInput) (PlaceResult, error)
}

type service struct {
	client Client
}

func NewService(c Service) *service {
	return &service{
		client: c,
	}
}
