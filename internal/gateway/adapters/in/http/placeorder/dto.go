package placeorder

type item struct {
	ID       int64  `json:"id"`
	Quantity int32  `json:"quantity"`
	Comment  string `json:"comment"`
}

type delivery struct {
	Type    string `json:"type"`
	Address string `json:"address,omitempty"`
	Comment string `json:"comment"`
}

type Request struct {
	UserID       int64  `json:"user_id"`
	RestaurantID int64  `json:"restaurant_id"`
	Items        []item `json:"items"`
	Delivery     delivery
}

type Response struct {
	OrderID int64  `json:"order_id"`
	Status  string `json:"status"`
}
