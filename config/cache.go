package config

type RedisConfig struct {
	RedisServer string `envconfig:"REDIS_SERVER" default:"localhost:6379"`
}
