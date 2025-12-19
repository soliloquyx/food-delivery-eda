package order

import (
	"context"
)

func (s *service) PlaceOrder(ctx context.Context, in PlaceOrderInput) (PlaceOrderResult, error) {
	result, err := s.client.PlaceOrder(ctx, in)
	if err != nil {
		return PlaceOrderResult{}, nil
	}

	return result, nil
}
