package order

import (
	"context"

	"github.com/google/uuid"
	orderapp "github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
)

type Client struct {
	svc orderv1.OrderServiceClient
}

func (c *Client) PlaceOrder(ctx context.Context, in orderapp.PlaceInput) (orderapp.PlaceResult, error) {
	req := &orderv1.PlaceOrderRequest{
		UserId:       in.UserID.String(),
		RestaurantId: in.RestaurantID.String(),
		Items:        itemsToProto(in.Items),
		Delivery:     deliveryToProto(in.Delivery),
	}

	resp, err := c.svc.PlaceOrder(ctx, req)
	if err != nil {
		return orderapp.PlaceResult{}, err
	}

	orderID, err := uuid.Parse(resp.GetOrderId())
	if err != nil {
		return orderapp.PlaceResult{}, err
	}

	return orderapp.PlaceResult{
		OrderID:   orderID,
		Status:    statusFromProto(resp.GetStatus()),
		CreatedAt: resp.CreatedAt.AsTime(),
	}, nil
}
