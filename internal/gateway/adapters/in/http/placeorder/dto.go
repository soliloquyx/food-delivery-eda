package placeorder

import "time"

type item struct {
	ID       string `json:"id"`
	Quantity int32  `json:"quantity"`
	Comment  string `json:"comment"`
}

type delivery struct {
	Type    string `json:"type"`
	Address string `json:"address,omitempty"`
	Comment string `json:"comment"`
}

type Request struct {
	UserID       string `json:"user_id"`
	RestaurantID string `json:"restaurant_id"`
	Items        []item `json:"items"`
	Delivery     delivery
}

type Response struct {
	OrderID   string    `json:"order_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
