package http

import "github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"

type handler struct {
	order order.Service
}

func NewHandler(o order.Service) *handler {
	return &handler{
		order: o,
	}
}
