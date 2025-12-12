package order

import (
	orderport "github.com/soliloquyx/food-delivery-eda/internal/gateway/ports/order"
)

type service struct {
	client orderport.Client
}

func NewService(c orderport.Client) *service {
	return &service{
		client: c,
	}
}
