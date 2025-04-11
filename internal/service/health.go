package service

import (
	"go-microservice/internal/shared"
	"go-microservice/internal/shared/dto"
)

type (
	Service interface {
		HealthCheck() (*dto.HealthCheckResponse, error)
	}

	implPlatform struct {
		deps shared.Deps
	}
)

func NewPlatformService(deps shared.Deps) (s Service, err error) {
	return &implPlatform{deps: deps}, nil
}

func (h *implPlatform) HealthCheck() (*dto.HealthCheckResponse, error) {
	resp := dto.HealthCheckResponse{
		Status: "OK",
	}
	return &resp, nil
}
