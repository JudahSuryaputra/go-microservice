package service

import (
	"go-microservice/internal/service/platform"
	"go-microservice/internal/service/user"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In

		PlatformService platform.Service
		UserService     user.Service
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(platform.NewPlatformService); err != nil {
		return err
	}
	if err := container.Provide(user.NewUserService); err != nil {
		return err
	}
	return nil
}
