package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Configuration struct {
	DbConfig
	RedisConfig
	RateLimitConfig

	AppMode string `envconfig:"APP_MODE" default:"DEBUG"`
}

func New() (*Configuration, error) {
	var cfg Configuration
	if err := LoadConfig(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func LoadConfig(target interface{}) error {
	var (
		filename = os.Getenv("CONFIG_FILE")
	)

	if filename == "" {
		filename = ".env"
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := envconfig.Process("", target); err != nil {
			return err
		}
		return nil
	}

	if err := godotenv.Load(filename); err != nil {
		return err
	}

	if err := envconfig.Process("", target); err != nil {
		return err
	}

	return nil
}
