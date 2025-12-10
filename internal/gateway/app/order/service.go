package order

import (
	orderout "github.com/soliloquyx/food-delivery-eda/internal/gateway/adapters/out/order"
	orderport "github.com/soliloquyx/food-delivery-eda/internal/gateway/ports/order"
)

type service struct {
	client orderout.Client
}

func NewService(c orderport.Client) *service {
	return &service{}
}
