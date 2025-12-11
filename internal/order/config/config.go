package config

import "github.com/caarlos0/env/v11"

type config struct {
	SvcName  string `env:"SERVICE_NAME"`
	GRPCAddr string `env:"GRPC_ADDR"`
}

func FromEnv() (config, error) {
	var c config
	err := env.Parse(&c)

	return c, err
}
