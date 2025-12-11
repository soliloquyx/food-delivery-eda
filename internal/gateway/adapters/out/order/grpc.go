package order

import (
	orderport "github.com/soliloquyx/food-delivery-eda/internal/gateway/ports/order"
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(addr string) (orderport.Service, func(), error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}

	c := &Client{
		svc: orderv1.NewOrderServiceClient(conn),
	}

	cleanup := func() {
		_ = conn.Close()
	}

	return c, cleanup, nil
}
