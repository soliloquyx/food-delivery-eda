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

func itemsToProto(items []orderapp.Item) []*orderv1.Item {
	var protoItems []*orderv1.Item
	for _, it := range items {
		protoItems = append(protoItems, &orderv1.Item{
			Id:       it.ID.String(),
			Quantity: it.Quantity,
			Comment:  it.Comment,
		})
	}

	return protoItems
}

func typeToProto(t orderapp.DeliveryType) orderv1.DeliveryType {
	switch t {
	case orderapp.DeliveryTypeDelivery:
		return orderv1.DeliveryType_DELIVERY_TYPE_DELIVERY
	case orderapp.DeliveryTypePickup:
		return orderv1.DeliveryType_DELIVERY_TYPE_PICKUP
	default:
		return orderv1.DeliveryType_DELIVERY_TYPE_UNSPECIFIED
	}
}

func deliveryToProto(d orderapp.Delivery) *orderv1.Delivery {
	return &orderv1.Delivery{
		Type:    typeToProto(d.Type),
		Address: d.Address,
		Comment: d.Comment,
	}
}
