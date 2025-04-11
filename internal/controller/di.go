package controller

import (
	"github.com/labstack/echo/v4"
	"go-microservice/internal/controller/data"
	"go-microservice/internal/shared"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In
		Deps               shared.Deps
		PlatformController *PlatformController
		DataController     *data.Controller
	}
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewPlatformController); err != nil {
		return err
	}
	if err := container.Provide(data.NewDataController); err != nil {
		return err
	}

	return nil
}

func (h *Holder) SetupRoutes(app *echo.Echo) {
	// check app health
	app.GET("/health", h.PlatformController.CheckSelf)

	/*
		v1 := app.Group("/v1")
		{
			v1.GET("/data")
		}
	*/
}
