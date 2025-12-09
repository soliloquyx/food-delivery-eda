package placeorder

import "github.com/soliloquyx/food-delivery-eda/internal/gateway/ports/order"

func ToInput(body Request) order.PlaceInput {
	var inputItems []order.Item
	for _, it := range body.Items {
		inputItems = append(inputItems, order.Item{
			ID:       it.ID,
			Quantity: it.Quantity,
			Comment:  it.Comment,
		})
	}

	return order.PlaceInput{
		UserID:       body.UserID,
		RestaurantID: body.RestaurantID,
		Items:        inputItems,
		Delivery: order.Delivery{
			Type:    body.Delivery.Type,
			Address: body.Delivery.Address,
			Comment: body.Delivery.Comment,
		},
	}
}
