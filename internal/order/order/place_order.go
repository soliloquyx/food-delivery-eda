package order

import (
	"context"

	"github.com/google/uuid"
)

func (s *service) PlaceOrder(ctx context.Context, in PlaceOrderInput) (PlaceOrderResult, error) {
	orderID, err := uuid.NewV7()
	if err != nil {
		return PlaceOrderResult{}, err
	}

	result, err := s.orderRepo.Create(ctx, orderID, in)
	if err != nil {
		return PlaceOrderResult{}, err
	}

	return result, nil
}
