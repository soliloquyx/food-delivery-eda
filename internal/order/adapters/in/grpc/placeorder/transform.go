package placeorder

import (
	"github.com/google/uuid"
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
	"github.com/soliloquyx/food-delivery-eda/internal/order/order"
)

func statusToProto(s order.Status) orderv1.Status {
	switch s {
	case order.StatusPending:
		return orderv1.Status_STATUS_PENDING
	case order.StatusConfirmed:
		return orderv1.Status_STATUS_CONFIRMED
	case order.StatusCancelled:
		return orderv1.Status_STATUS_CANCELLED
	default:
		return orderv1.Status_STATUS_UNSPECIFIED
	}
}

func ToInput(req *orderv1.PlaceOrderRequest) (order.PlaceOrderInput, error) {
	items := make([]order.OrderItem, len(req.Items))
	for _, it := range req.Items {
		itemID, err := uuid.Parse(it.Id)
		if err != nil {
			return order.PlaceOrderInput{}, err
		}

		items = append(items, order.OrderItem{
			ID:       itemID,
			Quantity: it.Quantity,
			Comment:  it.Comment,
		})
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return order.PlaceOrderInput{}, err
	}

	restaurantID, err := uuid.Parse(req.RestaurantId)
	if err != nil {
		return order.PlaceOrderInput{}, err
	}

	return order.PlaceOrderInput{
		UserID:       userID,
		RestaurantID: restaurantID,
		Items:        items,
		Delivery: order.Delivery{
			Type:    order.DeliveryType(req.Delivery.Type),
			Address: req.Delivery.Address,
			Comment: req.Delivery.Comment,
		},
	}, nil
}

func ToResponse(result order.PlaceOrderResult) *orderv1.PlaceOrderResponse {
	return &orderv1.PlaceOrderResponse{
		OrderId: result.OrderID.String(),
		Status:  statusToProto(result.Status),
	}
}
