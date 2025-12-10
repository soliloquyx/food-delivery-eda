package order

import (
	"context"

	orderport "github.com/soliloquyx/food-delivery-eda/internal/gateway/ports/order"
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
)

type Client struct {
	svc orderv1.OrderServiceClient
}

func (c *Client) Place(ctx context.Context, in orderport.PlaceInput) (orderport.PlaceResult, error) {
	req := &orderv1.PlaceOrderRequest{
		UserId:       in.UserID,
		RestaurantId: in.RestaurantID,
		Items:        itemsToProto(in.Items),
		Delivery:     deliveryToProto(in.Delivery),
	}

	resp, err := c.svc.PlaceOrder(ctx, req)
	if err != nil {
		return orderport.PlaceResult{}, err
	}

	return orderport.PlaceResult{
		OrderID: resp.GetOrderId(),
		Status:  statusFromProto(resp.GetStatus()),
	}, nil
}
