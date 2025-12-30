package placeorder

import "time"

type item struct {
	ItemID   string `json:"item_id"`
	Quantity int32  `json:"quantity"`
	Comment  string `json:"comment,omitempty"`
}

type delivery struct {
	Address string `json:"address"`
	Comment string `json:"comment,omitempty"`
}

type Request struct {
	UserID          string    `json:"user_id"`
	RestaurantID    string    `json:"restaurant_id"`
	Items           []item    `json:"items"`
	FulfillmentType string    `json:"fulfillment_type"`
	Delivery        *delivery `json:"delivery,omitempty"`
}

type Response struct {
	OrderID   string    `json:"order_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
