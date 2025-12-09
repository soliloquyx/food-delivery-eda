package order

import "context"

type Item struct {
	ID       int
	Quantity int
	Comment  string
}

type Delivery struct {
	Type    string
	Address string
	Comment string
}

type PlaceInput struct {
	UserID       int
	RestaurantID int
	Items        []Item
	Delivery     Delivery
}

type Service interface {
	Place(ctx context.Context, in PlaceInput) error
}
