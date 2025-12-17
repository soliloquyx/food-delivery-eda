package app

import "github.com/google/uuid"

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
	ID       uuid.UUID
	Quantity int32
	Comment  string
}

type Delivery struct {
	Type    DeliveryType
	Address string
	Comment string
}

type PlaceOrderInput struct {
	UserID       uuid.UUID
	RestaurantID uuid.UUID
	Items        []Item
	Delivery     Delivery
}

type PlaceOrderResult struct {
	OrderID uuid.UUID
	Status  Status
}
