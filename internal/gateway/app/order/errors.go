package order

import "errors"

var (
	ErrInvalidUUID            = errors.New("invalid uuid")
	ErrInvalidFulfillmentType = errors.New("invalid fulfillment type")
)
