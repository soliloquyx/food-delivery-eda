package placeorder

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
)

func toFulfillmentType(ft string) (order.FulfillmentType, error) {
	switch ft {
	case "delivery":
		return order.FulfillmentTypeDelivery, nil
	case "pickup":
		return order.FulfillmentTypePickup, nil
	default:
		return "", order.ErrInvalidFulfillmentType
	}
}

func ToInput(body Request) (order.PlaceOrderInput, error) {
	inputItems := make([]order.OrderItem, 0, len(body.Items))
	for _, it := range body.Items {
		itemID, err := uuid.Parse(it.ItemID)
		if err != nil {
			return order.PlaceOrderInput{}, fmt.Errorf("%w: %w", order.ErrInvalidUUID, err)
		}

		inputItems = append(inputItems, order.OrderItem{
			ItemID:   itemID,
			Quantity: it.Quantity,
			Comment:  it.Comment,
		})
	}

	userID, err := uuid.Parse(body.UserID)
	if err != nil {
		return order.PlaceOrderInput{}, fmt.Errorf("%w: %w", order.ErrInvalidUUID, err)
	}

	restaurantID, err := uuid.Parse(body.RestaurantID)
	if err != nil {
		return order.PlaceOrderInput{}, fmt.Errorf("%w: %w", order.ErrInvalidUUID, err)
	}

	ft, err := toFulfillmentType(body.FulfillmentType)
	if err != nil {
		return order.PlaceOrderInput{}, err
	}

	var delivery *order.Delivery
	if body.Delivery != nil {
		delivery = &order.Delivery{
			Address: body.Delivery.Address,
			Comment: body.Delivery.Comment,
		}
	}

	return order.PlaceOrderInput{
		UserID:          userID,
		RestaurantID:    restaurantID,
		Items:           inputItems,
		FulfillmentType: ft,
		Delivery:        delivery,
	}, nil
}

func ToResponse(result order.PlaceOrderResult) Response {
	return Response{
		OrderID:   result.OrderID.String(),
		Status:    string(result.Status),
		CreatedAt: result.CreatedAt,
	}
}
