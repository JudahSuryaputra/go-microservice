package data

import (
	"go-microservice/internal/service"
	"go-microservice/internal/shared"
)

type (
	Controller struct {
		deps     shared.Deps
		services service.Holder
	}
)

func NewDataController(deps shared.Deps, services service.Holder) (*Controller, error) {
	return &Controller{deps: deps, services: services}, nil
}
