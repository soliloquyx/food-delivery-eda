package order

import (
	"context"

	orderport "github.com/soliloquyx/food-delivery-eda/internal/gateway/ports/order"
)

func (s *service) Place(ctx context.Context, in orderport.PlaceInput) (orderport.PlaceResult, error) {
	result, err := s.client.Place(ctx, in)
	if err != nil {
		return orderport.PlaceResult{}, nil
	}

	return result, nil
}
