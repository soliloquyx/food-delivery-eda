package placeorder

import (
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
	"github.com/soliloquyx/food-delivery-eda/internal/order/ports"
)

func statusToProto(s ports.Status) orderv1.Status {
	switch s {
	case ports.StatusPending:
		return orderv1.Status_STATUS_PENDING
	case ports.StatusConfirmed:
		return orderv1.Status_STATUS_CONFIRMED
	case ports.StatusCancelled:
		return orderv1.Status_STATUS_CANCELLED
	default:
		return orderv1.Status_STATUS_UNSPECIFIED
	}
}

func ToInput(req *orderv1.PlaceOrderRequest) ports.PlaceOrderInput {
	items := make([]ports.Item, len(req.Items))
	for _, it := range items {
		items = append(items, ports.Item{
			ID:       it.ID,
			Quantity: it.Quantity,
			Comment:  it.Comment,
		})
	}

	return ports.PlaceOrderInput{
		UserID:       req.UserId,
		RestaurantID: req.RestaurantId,
		Items:        items,
		Delivery: ports.Delivery{
			Type:    ports.DeliveryType(req.Delivery.Type),
			Address: req.Delivery.Address,
			Comment: req.Delivery.Comment,
		},
	}
}

func ToResponse(result ports.PlaceOrderResult) *orderv1.PlaceOrderResponse {
	return &orderv1.PlaceOrderResponse{
		OrderId: result.OrderID,
		Status:  statusToProto(result.Status),
	}
}
