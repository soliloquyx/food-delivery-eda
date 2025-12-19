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

type FulfillmentType string

const (
	FulfillmentTypeUnknown  FulfillmentType = "unknown"
	FulfillmentTypeDelivery FulfillmentType = "delivery"
	FulfillmentTypePickup   FulfillmentType = "pickup"
)

type OrderItem struct {
	ItemID   uuid.UUID
	Quantity int32
	Comment  string
}

type Delivery struct {
	Address string
	Comment string
}

type PlaceOrderInput struct {
	UserID          uuid.UUID
	RestaurantID    uuid.UUID
	Items           []OrderItem
	FulfillmentType FulfillmentType
	Delivery        *Delivery
}

type PlaceOrderResult struct {
	OrderID   uuid.UUID
	Status    Status
	CreatedAt time.Time
}

type OrderRepo interface {
	Create(ctx context.Context, orderID uuid.UUID, in PlaceOrderInput) (PlaceOrderResult, error)
}
