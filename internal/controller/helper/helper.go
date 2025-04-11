package helper

import (
	"github.com/labstack/echo/v4"
	"go-microservice/internal/shared/dto"
	"net/http"
)

func SuccessResponse(c echo.Context, data interface{}) error {
	var (
		res = dto.Success{
			Code:    "200",
			Message: "success",
			Data:    data,
		}
	)

	return c.JSON(http.StatusOK, res)
}
