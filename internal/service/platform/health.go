package platform

import (
	"context"
	"go-microservice/internal/repository"
	"go-microservice/internal/shared"
	"go-microservice/internal/shared/dto"
)

type (
	Service interface {
		HealthCheck(ctx context.Context) (*dto.HealthCheckResponse, error)
	}

	implPlatform struct {
		deps shared.Deps

		repo repository.Holder
	}
)

func NewPlatformService(deps shared.Deps, repo repository.Holder) Service {
	return &implPlatform{deps: deps, repo: repo}
}

func (i *implPlatform) HealthCheck(ctx context.Context) (*dto.HealthCheckResponse, error) {
	var (
		redisStatus = "error"
	)

	_, err := i.repo.CacheRepository.CheckHealth(ctx)
	if err == nil {
		redisStatus = "OK"
	}

	resp := dto.HealthCheckResponse{
		Status:      "OK",
		RedisStatus: redisStatus,
	}
	return &resp, nil
}
