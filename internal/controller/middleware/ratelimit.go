package middleware

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

func (mw *implMiddleware) RateLimit(next echo.HandlerFunc) echo.HandlerFunc {
	limiter := rate.NewLimiter(rate.Limit(mw.deps.Config.RPS), mw.deps.Config.Burst)
	return func(ctx echo.Context) error {
		if !limiter.Allow() {
			return errors.New("Too Many Requests")
		}
		return next(ctx)
	}
}
