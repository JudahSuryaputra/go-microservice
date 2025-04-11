package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	DbConfig
	RedisConfig
}

func New() *Configuration {
	var cfg Configuration
	envconfig.Process("", &cfg)
	return &cfg
}
