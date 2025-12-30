package orderclient

import (
	orderapp "github.com/soliloquyx/food-delivery-eda/internal/gateway/app/order"
	orderv1 "github.com/soliloquyx/food-delivery-eda/internal/genproto/order/v1"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(addr string) (orderapp.Service, func(), error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
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
