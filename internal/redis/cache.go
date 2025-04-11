package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-microservice/config"
)

func InitRedis(cfg *config.Configuration) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisServer,
	})
	result, err := client.Ping(context.Background()).Result()
	fmt.Printf("Redis ping result: %v\n", result)
	if err != nil {
		panic(err)
	}
	return client
}
