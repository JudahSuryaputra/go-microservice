package user

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-microservice/internal/controller/helper"
	"go-microservice/internal/controller/helper/validation"
	"go-microservice/internal/service"
	"go-microservice/internal/shared"
	"go-microservice/internal/shared/dto"
)

type (
	Controller struct {
		deps     shared.Deps
		services service.Holder
	}
)

func NewUserController(deps shared.Deps, services service.Holder) (*Controller, error) {
	return &Controller{deps: deps, services: services}, nil
}

func (c *Controller) GetUserByID(ctx echo.Context) (err error) {
	var (
		rctx = ctx.Request().Context()
	)

	resp, err := c.services.UserService.GetUserByID(rctx, "userID")
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, resp)
}

func (c *Controller) CreateUser(ctx echo.Context) (err error) {
	var (
		rctx = ctx.Request().Context()
		req  = dto.CreateUserRequest{}
	)

	if err = json.NewDecoder(ctx.Request().Body).Decode(&req); err != nil {
		return helper.ErrorResponse(ctx, err)
	}

	if err = validation.Struct(req); err != nil {
		return helper.ErrorResponse(ctx, err)
	}

	resp, err := c.services.UserService.CreateUser(rctx, &req)
	if err != nil {
		return helper.ErrorResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, resp)
}
