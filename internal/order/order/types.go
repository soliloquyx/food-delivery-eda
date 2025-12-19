package order

import (
	"context"
	"time"

	"github.com/google/uuid"
)

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

type OrderItem struct {
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
	Items        []OrderItem
	Delivery     Delivery
}

type PlaceOrderResult struct {
	OrderID   uuid.UUID
	Status    Status
	CreatedAt time.Time
}

type OrderRepo interface {
	Create(ctx context.Context, orderID uuid.UUID, in PlaceOrderInput) (PlaceOrderResult, error)
}
