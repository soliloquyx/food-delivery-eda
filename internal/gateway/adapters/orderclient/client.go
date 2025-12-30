package orderclient

import (
	"context"

	"github.com/google/uuid"
	orderapp "github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
)

type Client struct {
	svc orderv1.OrderServiceClient
}

func (c *Client) PlaceOrder(ctx context.Context, in orderapp.PlaceOrderInput) (orderapp.PlaceOrderResult, error) {
	req := &orderv1.PlaceOrderRequest{
		UserId:          in.UserID.String(),
		RestaurantId:    in.RestaurantID.String(),
		Items:           itemsToProto(in.Items),
		FulfillmentType: fulfillmentTypeToProto(in.FulfillmentType),
	}

	if in.Delivery != nil {
		req.Delivery = deliveryToProto(*in.Delivery)
	}

	resp, err := c.svc.PlaceOrder(ctx, req)
	if err != nil {
		return orderapp.PlaceOrderResult{}, err
	}

	orderID, err := uuid.Parse(resp.GetOrderId())
	if err != nil {
		return orderapp.PlaceOrderResult{}, err
	}

	return orderapp.PlaceOrderResult{
		OrderID:   orderID,
		Status:    statusFromProto(resp.GetStatus()),
		CreatedAt: resp.CreatedAt.AsTime(),
	}, nil
}
