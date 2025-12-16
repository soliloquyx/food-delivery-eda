package placeorder

import (
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

func ToInput(req *orderv1.PlaceOrderRequest) app.PlaceOrderInput {
	items := make([]app.Item, len(req.Items))
	for _, it := range items {
		items = append(items, app.Item{
			ID:       it.ID,
			Quantity: it.Quantity,
			Comment:  it.Comment,
		})
	}

	return app.PlaceOrderInput{
		UserID:       req.UserId,
		RestaurantID: req.RestaurantId,
		Items:        items,
		Delivery: app.Delivery{
			Type:    app.DeliveryType(req.Delivery.Type),
			Address: req.Delivery.Address,
			Comment: req.Delivery.Comment,
		},
	}
}

func ToResponse(result app.PlaceOrderResult) *orderv1.PlaceOrderResponse {
	return &orderv1.PlaceOrderResponse{
		OrderId: result.OrderID,
		Status:  statusToProto(result.Status),
	}
}
