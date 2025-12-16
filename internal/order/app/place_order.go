package app

import "context"

func (s *service) PlaceOrder(ctx context.Context, in PlaceOrderInput) (PlaceOrderResult, error) {
	return PlaceOrderResult{
		OrderID: 1,
		Status:  StatusConfirmed,
	}, nil
}
