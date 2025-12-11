package ports

import "context"

type Status string

const (
	StatusUnknown   Status = "unknown"
	StatusPending   Status = "pending"
	StatusConfirmed Status = "confirmed"
	StatusCancelled Status = "cancelled"
)

type DeliveryType string

const (
	DeliveryTypeUnknown  DeliveryType = "unknown"
	DeliveryTypeDelivery DeliveryType = "delivery"
	DeliveryTypePickup   DeliveryType = "pickup"
)

type Item struct {
	ID       int64
	Quantity int32
	Comment  string
}

type Delivery struct {
	Type    DeliveryType
	Address string
	Comment string
}

type PlaceOrderInput struct {
	UserID       int64
	RestaurantID int64
	Items        []Item
	Delivery     Delivery
}

type PlaceOrderResult struct {
	OrderID int64
	Status  Status
}

type Service interface {
	PlaceOrder(ctx context.Context, in PlaceOrderInput) (PlaceOrderResult, error)
}
