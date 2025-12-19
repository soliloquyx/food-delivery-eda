package placeorder

import (
	"github.com/google/uuid"
	"github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
)

func ToInput(body Request) (order.PlaceInput, error) {
	var inputItems []order.OrderItem
	for _, it := range body.Items {
		itemID, err := uuid.Parse(it.ID)
		if err != nil {
			return order.PlaceInput{}, err
		}

		inputItems = append(inputItems, order.OrderItem{
			ID:       itemID,
			Quantity: it.Quantity,
			Comment:  it.Comment,
		})
	}

	userID, err := uuid.Parse(body.UserID)
	if err != nil {
		return order.PlaceInput{}, err
	}

	restaurantID, err := uuid.Parse(body.RestaurantID)
	if err != nil {
		return order.PlaceInput{}, err
	}

	return order.PlaceInput{
		UserID:       userID,
		RestaurantID: restaurantID,
		Items:        inputItems,
		Delivery: order.Delivery{
			Type:    order.DeliveryType(body.Delivery.Type),
			Address: body.Delivery.Address,
			Comment: body.Delivery.Comment,
		},
	}, nil
}

func ToResponse(result order.PlaceResult) Response {
	return Response{
		OrderID:   result.OrderID.String(),
		Status:    string(result.Status),
		CreatedAt: result.CreatedAt,
	}
}
