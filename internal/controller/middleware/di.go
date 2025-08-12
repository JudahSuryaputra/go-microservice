package middleware

import (
	"github.com/labstack/echo/v4"
	"go-microservice/internal/repository"
	"go-microservice/internal/shared"
	"go.uber.org/dig"
)

type Middleware interface {
	RateLimit(next echo.HandlerFunc) echo.HandlerFunc
}

type implMiddleware struct {
	repo repository.Holder
	deps shared.Deps
}

func NewMiddlewares(repo repository.Holder, deps shared.Deps) Middleware {
	return &implMiddleware{repo: repo, deps: deps}
}

func Register(container *dig.Container) error {
	return container.Provide(NewMiddlewares)
}
