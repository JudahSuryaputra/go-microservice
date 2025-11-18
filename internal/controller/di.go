package controller

import (
	"go-microservice/internal/controller/middleware"
	"go-microservice/internal/controller/user"
	"go-microservice/internal/shared"

	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	Holder struct {
		dig.In
		Deps               shared.Deps
		InternalMiddleware middleware.Middleware
		PlatformController *PlatformController
		UserController     *user.Controller
	}
)

func Register(container *dig.Container) error {
	if err := middleware.Register(container); err != nil {
		return err
	}
	if err := container.Provide(NewPlatformController); err != nil {
		return err
	}
	if err := container.Provide(user.NewUserController); err != nil {
		return err
	}

	return nil
}

func (h *Holder) SetupRoutes(app *echo.Echo) {
	// check app health
	app.Use(h.InternalMiddleware.Logging())
	app.Use(h.InternalMiddleware.RateLimit)
	app.GET("/health", h.PlatformController.CheckSelf)

	v1 := app.Group("/v1")
	v1.Use(h.InternalMiddleware.RateLimit)
	{
		v1.GET("/user/:id", h.UserController.GetUserByID)
	}

}
