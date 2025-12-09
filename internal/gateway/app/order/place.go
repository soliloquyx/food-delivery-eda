package order

import (
	"context"

	orderport "github.com/soliloquyx/food-delivery-eda/internal/gateway/ports/order"
)

func (s *service) Place(ctx context.Context, in orderport.PlaceInput) error {
	return nil
}
