package log

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-todo/internal/log"
)

func EchoLogger(skipper middleware.Skipper) echo.MiddlewareFunc {
	logFormat := `
==================================================================
URI: ${uri}
METHOD: ${method}
STATUS: ${status}
TIME: ${time_rfc3339_nano}
LATENCY: ${latency_human}
AGENT: ${user_agent}
==================================================================
		` + "\n"

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Skipper: skipper,
		Format:  logFormat,
	})
}

func BodyDumpLogger(skipper middleware.Skipper) echo.MiddlewareFunc {
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: skipper,
		Handler: func(ec echo.Context, req []byte, res []byte) {
			headers := "[\n"
			for header, value := range ec.Request().Header {
				headers += fmt.Sprintf("%s: %s\n", header, value)
			}
			headers += "]"

			log.Logger.Infof(
				"\n==================================================================\n"+
					"HEADERS: %s\n"+
					"REQUEST: %s\n"+
					"RESPONSE: %s"+
					"==================================================================",
				headers,
				string(req),
				string(res),
			)
		},
	})
}
