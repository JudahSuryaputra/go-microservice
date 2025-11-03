package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-microservice/internal/shared/utils"
	"io"
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

			// - Writer
			w := io.MultiWriter(ctx.Response().Writer, resBody)
			writer := &utils.BodyDumpResponseWriter{
				Writer:         w,
				ResponseWriter: ctx.Response().Writer,
			}
			// Replace Echo's writer with our wrapper
			ctx.Response().Writer = writer
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
		reqHeader, _ = json.Marshal(ctx.Request().Header)
		reqBody      []byte
	)

	// ignore health check log
	//if ctx.Request().URL.String() == "/health" {
	//	return
	//}

	// read body (and restore for handlers)
	if ctx.Request().Body != nil {
		b, _ := io.ReadAll(ctx.Request().Body)
		reqBody = b
		ctx.Request().Body = io.NopCloser(bytes.NewBuffer(b))
	}

	log.WithContext(ctx.Request().Context()).WithFields(log.Fields{
		"headers": tryParseJSON(reqHeader),
		"request": tryParseJSON(reqBody),
	}).Info(fmt.Sprintf("Incoming: %s", ctx.Request().URL.String()))
}

// logResponse read response, and put in log
func (mw *implMiddleware) logResponse(ctx echo.Context, resBody *bytes.Buffer) {
	// ignore health check log
	//if ctx.Request().URL.String() == "/health" {
	//	return
	//}

	requestTime := ctx.Get(utils.ReqTimeStart).(time.Time)
	respHeader, _ := json.Marshal(ctx.Response().Header())

	log.WithContext(ctx.Request().Context()).WithFields(log.Fields{
		"status":        ctx.Response().Status,
		"response_time": fmt.Sprint(time.Since(requestTime)),
		"headers":       tryParseJSON(respHeader),
		"response":      tryParseJSON(resBody.Bytes()),
	}).Info(fmt.Sprintf("Outgoing: %s", ctx.Request().URL.String()))
}

func tryParseJSON(b []byte) interface{} {
	if len(b) == 0 {
		return nil
	}
	var v interface{}
	if err := json.Unmarshal(b, &v); err == nil {
		return v // ✅ return as map[string]interface{}
	}
	return string(b)
}
