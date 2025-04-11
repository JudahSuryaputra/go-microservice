package service

import (
	"go-microservice/internal/service/data"
	"go-microservice/internal/service/platform"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In

		PlatformService platform.Service
		DataService     data.Service
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(platform.NewPlatformService); err != nil {
		return err
	}
	if err := container.Provide(data.NewDataService); err != nil {
		return err
	}
	return nil
}
