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

func toFulfillmentType(ft orderv1.FulfillmentType) order.FulfillmentType {
	switch ft {
	case orderv1.FulfillmentType_FULFILLMENT_TYPE_DELIVERY:
		return order.FulfillmentTypeDelivery
	case orderv1.FulfillmentType_FULFILLMENT_TYPE_PICKUP:
		return order.FulfillmentTypePickup
	default:
		return ""
	}
}

func ToInput(req *orderv1.PlaceOrderRequest) (order.PlaceOrderInput, error) {
	items := make([]order.OrderItem, 0, len(req.Items))
	for _, it := range req.Items {
		itemID, err := uuid.Parse(it.ItemId)
		if err != nil {
			return order.PlaceOrderInput{}, err
		}

		items = append(items, order.OrderItem{
			ItemID:   itemID,
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

	in := order.PlaceOrderInput{
		UserID:          userID,
		RestaurantID:    restaurantID,
		Items:           items,
		FulfillmentType: toFulfillmentType(req.FulfillmentType),
	}

	if d := req.GetDelivery(); d != nil {
		in.Delivery = &order.Delivery{
			Address: d.Address,
			Comment: d.Comment,
		}
	}

	return in, nil
}

func ToResponse(result order.PlaceOrderResult) *orderv1.PlaceOrderResponse {
	return &orderv1.PlaceOrderResponse{
		OrderId: result.OrderID.String(),
		Status:  statusToProto(result.Status),
	}
}
