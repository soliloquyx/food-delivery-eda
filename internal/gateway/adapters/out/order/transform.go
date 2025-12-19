package order

import (
	orderapp "github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
)

func statusFromProto(s orderv1.Status) orderapp.Status {
	switch s {
	case orderv1.Status_STATUS_PENDING:
		return orderapp.StatusPending
	case orderv1.Status_STATUS_CONFIRMED:
		return orderapp.StatusConfirmed
	case orderv1.Status_STATUS_CANCELLED:
		return orderapp.StatusCancelled
	default:
		return orderapp.StatusUnknown
	}
}

func itemsToProto(items []orderapp.OrderItem) []*orderv1.OrderItem {
	var protoItems []*orderv1.OrderItem
	for _, it := range items {
		protoItems = append(protoItems, &orderv1.OrderItem{
			ItemId:   it.ItemID.String(),
			Quantity: it.Quantity,
			Comment:  it.Comment,
		})
	}

	return protoItems
}

func fulfillmentTypeToProto(t orderapp.FulfillmentType) orderv1.FulfillmentType {
	switch t {
	case orderapp.FulfillmentTypeDelivery:
		return orderv1.FulfillmentType_FULFILLMENT_TYPE_FULFILLMENT
	case orderapp.FulfillmentTypePickup:
		return orderv1.FulfillmentType_FULFILLMENT_TYPE_PICKUP
	default:
		return orderv1.FulfillmentType_FULFILLMENT_TYPE_UNSPECIFIED
	}
}

func deliveryToProto(d orderapp.Delivery) *orderv1.Delivery {
	return &orderv1.Delivery{
		Address: d.Address,
		Comment: d.Comment,
	}
}
