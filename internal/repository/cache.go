package repository

import (
	"context"
	"fmt"
	"go-microservice/internal/shared"
)

type (
	CacheRepository interface {
		CheckHealth(ctx context.Context) (string, error)
		Save(ctx context.Context, key string, value string) error
		Get(ctx context.Context, key string) (string, error)
	}

	cacheRepository struct {
		deps shared.Deps
	}
)

func NewCacheRepository(deps shared.Deps) CacheRepository {
	return &cacheRepository{deps: deps}
}

func (r *cacheRepository) CheckHealth(ctx context.Context) (string, error) {
	result, err := r.deps.RedisClient.Ping(ctx).Result()
	fmt.Printf("Redis ping result: %v\n", result)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (r *cacheRepository) Save(ctx context.Context, key string, value string) error {
	r.deps.RedisClient.Set(ctx, key, value, 36000000)
	return nil
}

func (r *cacheRepository) Get(ctx context.Context, key string) (string, error) {
	return r.deps.RedisClient.Get(ctx, key).Result()
}
