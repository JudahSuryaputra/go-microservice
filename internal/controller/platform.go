package controller

import (
	"github.com/labstack/echo/v4"
	"go-microservice/internal/controller/helper"
	"go-microservice/internal/service"
	"go-microservice/internal/shared"
)

type (
	PlatformController struct {
		deps     shared.Deps
		services service.Holder
	}
)

func NewPlatformController(services service.Holder, deps shared.Deps) (*PlatformController, error) {
	return &PlatformController{services: services, deps: deps}, nil
}

func (c *PlatformController) CheckSelf(ctx echo.Context) error {
	resp, err := c.services.PlatformService.HealthCheck(ctx.Request().Context())
	if err != nil {
		return err
	}

	return helper.SuccessResponse(ctx, resp)
}
