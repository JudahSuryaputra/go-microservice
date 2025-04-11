package service

import (
	"go-microservice/internal/service/data"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In

		PlatformService Service
		DataService     data.Service
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewPlatformService); err != nil {
		return err
	}
	if err := container.Provide(data.NewDataService); err != nil {
		return err
	}
	return nil
}
