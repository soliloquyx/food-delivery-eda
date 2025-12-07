package httpapi

import "github.com/soliloquyx/food-delivery-eda/internal/gateway/ports/order"

type handler struct {
	Order order.Service
}

func NewHandler(o order.Service) *handler {
	return &handler{
		Order: o,
	}
}
