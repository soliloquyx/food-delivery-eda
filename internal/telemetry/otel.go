package telemetry

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

type Config struct {
	SvcName  string
	Endpoint string
}

func InitOTLP(ctx context.Context, cfg Config) (shutdown func(context.Context) error, err error) {
	exp, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithEndpoint(cfg.Endpoint),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	res := resource.NewWithAttributes(
		"",
		attribute.String("service.name", cfg.SvcName),
	)

	tp := trace.NewTracerProvider(
		trace.WithBatcher(
			exp,
			trace.WithBatchTimeout(2*time.Second),
		),
		trace.WithResource(res),
	)

	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return tp.Shutdown, nil
}
