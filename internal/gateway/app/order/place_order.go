package order

import (
	"context"
)

func (s *service) PlaceOrder(ctx context.Context, in PlaceOrderInput) (PlaceOrderResult, error) {
	return s.client.PlaceOrder(ctx, in)
}
