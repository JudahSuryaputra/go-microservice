package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-microservice/internal/shared"
)

type (
	CacheRepository interface {
		SetValue(client *redis.Client, key string, value string) error
		GetValue(client *redis.Client, key string) (string, error)
	}

	implCacheRepository struct {
		deps shared.Deps
	}
)

func NewCacheRepository(deps shared.Deps) CacheRepository {
	return &implCacheRepository{deps: deps}
}

func (c *implCacheRepository) SetValue(client *redis.Client, key string, value string) error {
	ctx := context.Background()
	return client.Set(ctx, key, value, 0).Err()
}

func (c *implCacheRepository) GetValue(client *redis.Client, key string) (string, error) {
	ctx := context.Background()
	return client.Get(ctx, key).Result()
}
