package config

import "github.com/caarlos0/env/v11"

type config struct {
	SvcName       string `env:"SERVICE_NAME"`
	HTTPAddr      string `env:"HTTP_ADDR"`
	OrderGRPCAddr string `env:"ORDER_GRPC_ADDR"`
	OTLPEndpoint  string `env:"OTEL_EXPORTER_OTLP_ENDPOINT"`
}

func FromEnv() (config, error) {
	var c config
	err := env.Parse(&c)

	return c, err
}
