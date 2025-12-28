package order

func (it *OrderItem) validate() error {
	if it.Quantity <= 0 {
		return ErrInvalidItemQuantity
	}

	return nil
}

func (in *PlaceOrderInput) validate() error {
	switch in.FulfillmentType {
	case FulfillmentTypePickup, FulfillmentTypeDelivery:
	default:
		return ErrInvalidFulfillmentType
	}

	for _, it := range in.Items {
		if err := it.validate(); err != nil {
			return err
		}
	}

	return nil
}
