package order

import (
	"context"

	orderapp "github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
)

type Client struct {
	svc orderv1.OrderServiceClient
}

func (c *Client) Place(ctx context.Context, in orderapp.PlaceInput) (orderapp.PlaceResult, error) {
	req := &orderv1.PlaceOrderRequest{
		UserId:       in.UserID,
		RestaurantId: in.RestaurantID,
		Items:        itemsToProto(in.Items),
		Delivery:     deliveryToProto(in.Delivery),
	}

	resp, err := c.svc.PlaceOrder(ctx, req)
	if err != nil {
		return orderapp.PlaceResult{}, err
	}

	return orderapp.PlaceResult{
		OrderID: resp.GetOrderId(),
		Status:  statusFromProto(resp.GetStatus()),
	}, nil
}
