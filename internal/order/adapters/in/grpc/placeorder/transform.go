package placeorder

import (
	"github.com/google/uuid"
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
	"github.com/soliloquyx/food-delivery-eda/internal/order/app"
)

func statusToProto(s app.Status) orderv1.Status {
	switch s {
	case app.StatusPending:
		return orderv1.Status_STATUS_PENDING
	case app.StatusConfirmed:
		return orderv1.Status_STATUS_CONFIRMED
	case app.StatusCancelled:
		return orderv1.Status_STATUS_CANCELLED
	default:
		return orderv1.Status_STATUS_UNSPECIFIED
	}
}

func ToInput(req *orderv1.PlaceOrderRequest) (app.PlaceOrderInput, error) {
	items := make([]app.OrderItem, len(req.Items))
	for _, it := range items {
		items = append(items, app.OrderItem{
			ID:       it.ID,
			Quantity: it.Quantity,
			Comment:  it.Comment,
		})
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return app.PlaceOrderInput{}, err
	}

	restaurantID, err := uuid.Parse(req.RestaurantId)
	if err != nil {
		return app.PlaceOrderInput{}, err
	}

	return app.PlaceOrderInput{
		UserID:       userID,
		RestaurantID: restaurantID,
		Items:        items,
		Delivery: app.Delivery{
			Type:    app.DeliveryType(req.Delivery.Type),
			Address: req.Delivery.Address,
			Comment: req.Delivery.Comment,
		},
	}, nil
}

func ToResponse(result app.PlaceOrderResult) *orderv1.PlaceOrderResponse {
	return &orderv1.PlaceOrderResponse{
		OrderId: result.OrderID.String(),
		Status:  statusToProto(result.Status),
	}
}
