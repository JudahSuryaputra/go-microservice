package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-microservice/internal/shared/utils"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func (mw *implMiddleware) Logging() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			var (
				reqTimeStart = time.Now()
				resBody      = new(bytes.Buffer)
			)

			ctx.Set(utils.ReqTimeStart, reqTimeStart)
			ctx.Set(utils.RcCode, "2000000")

			// - middleware body dump start
			// - Request
			log.SetFormatter(&log.JSONFormatter{})
			mw.logRequest(ctx)

			// - Response
			// Wrap the original writer
			bdw := &utils.BodyDumpResponseWriter{
				ResponseWriter: ctx.Response().Writer,
				Body:           resBody,
			}

			// Replace Echo's writer with our wrapper
			ctx.Response().Writer = bdw
			// - middleware body dump end

			if err := next(ctx); err != nil {
				ctx.Error(err)
			}

			mw.logResponse(ctx, resBody)
			log.SetOutput(os.Stdout)

			return nil
		}
	}
}

// logRequest read request, and put in log
func (mw *implMiddleware) logRequest(ctx echo.Context) {
	var (
		headers = ctx.Request().Header
		reqBody []byte
	)

	// ignore health check log
	//if ctx.Request().URL.String() == "/health" {
	//	return
	//}

	log.WithContext(ctx.Request().Context()).Info(
		fmt.Sprintf("Incoming: %s \nHeaders: %s", ctx.Request().URL.String(), headers))
	log.WithFields(log.Fields{
		"headers": headers,
		"request": string(reqBody),
	})
}

// logResponse read response, and put in log
func (mw *implMiddleware) logResponse(ctx echo.Context, resBody *bytes.Buffer) {
	// ignore health check log
	//if ctx.Request().URL.String() == "/health" {
	//	return
	//}

	requestTime := ctx.Get(utils.ReqTimeStart).(time.Time)
	respHeader, _ := json.Marshal(ctx.Response().Header())

	log.WithContext(ctx.Request().Context()).Info(
		fmt.Sprintf("Outgoing: %s, to: %s", ctx.Request().URL.String(), respHeader))
	log.WithFields(log.Fields{
		"status":        ctx.Response().Status,
		"response_time": fmt.Sprint(time.Since(requestTime)),
		"headers":       string(respHeader),
		"request":       string(resBody.Bytes()),
	})
}
