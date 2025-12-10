package order

import (
	orderport "github.com/soliloquyx/food-delivery-eda/internal/gateway/ports/order"
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
)

func statusFromProto(s orderv1.Status) orderport.Status {
	switch s {
	case orderv1.Status_STATUS_PENDING:
		return orderport.StatusPending
	case orderv1.Status_STATUS_CONFIRMED:
		return orderport.StatusConfirmed
	case orderv1.Status_STATUS_CANCELLED:
		return orderport.StatusCancelled
	default:
		return orderport.StatusUnknown
	}
}

func itemsToProto(items []orderport.Item) []*orderv1.Item {
	var protoItems []*orderv1.Item
	for _, it := range items {
		protoItems = append(protoItems, &orderv1.Item{
			Id:       it.ID,
			Quantity: it.Quantity,
			Comment:  it.Comment,
		})
	}

	return protoItems
}

func typeToProto(t orderport.DeliveryType) orderv1.DeliveryType {
	switch t {
	case orderport.DeliveryTypeDelivery:
		return orderv1.DeliveryType_DELIVERY_TYPE_DELIVERY
	case orderport.DeliveryTypePickup:
		return orderv1.DeliveryType_DELIVERY_TYPE_PICKUP
	default:
		return orderv1.DeliveryType_DELIVERY_TYPE_UNSPECIFIED
	}
}

func deliveryToProto(d orderport.Delivery) *orderv1.Delivery {
	return &orderv1.Delivery{
		Type:    typeToProto(d.Type),
		Address: d.Address,
		Comment: d.Comment,
	}
}
