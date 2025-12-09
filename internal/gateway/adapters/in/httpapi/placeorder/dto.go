package placeorder

type item struct {
	ID       int    `json:"id"`
	Quantity int    `json:"quantity"`
	Comment  string `json:"comment"`
}

type delivery struct {
	Type    string `json:"type"`
	Address string `json:"address,omitempty"`
	Comment string `json:"comment"`
}

type Request struct {
	UserID       int    `json:"user_id"`
	RestaurantID int    `json:"restaurant_id"`
	Items        []item `json:"items"`
	Delivery     delivery
}
