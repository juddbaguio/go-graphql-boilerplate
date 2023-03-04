package config

import (
	"github.com/caarlos0/env/v7"
)

type Container struct {
	Mongo Mongo `envPrefix:"MONGO_"`
	// any other infra configs
}

func SetupConfig() (*Container, error) {
	var config Container

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
