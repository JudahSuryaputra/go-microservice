package di

import (
	"go-microservice/config"
	"go-microservice/internal/controller"
	"go-microservice/internal/repository"
	"go-microservice/internal/service"
	"go-microservice/internal/shared"
	"go.uber.org/dig"
)

var (
	Container = dig.New()
)

func init() {
	if err := Container.Provide(config.New); err != nil {
		panic(err)
	}

	if err := Container.Provide(NewORM); err != nil {
		panic(err)
	}

	if err := controller.Register(Container); err != nil {
		panic(err)
	}
	if err := service.Register(Container); err != nil {
		panic(err)
	}
	if err := repository.Register(Container); err != nil {
		panic(err)
	}
	if err := shared.Register(Container); err != nil {
		panic(err)
	}
}
