package order

import (
	"context"
)

func (s *service) PlaceOrder(ctx context.Context, in PlaceInput) (PlaceResult, error) {
	result, err := s.client.PlaceOrder(ctx, in)
	if err != nil {
		return PlaceResult{}, nil
	}

	return result, nil
}
