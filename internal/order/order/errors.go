package order

import "errors"

var (
	ErrInvalidFulfillmentType = errors.New("invalid fulfillment type")
	ErrInvalidItemQuantity    = errors.New("invalid order item quantity")
)
