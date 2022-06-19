package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-todo/server/config"
	"go-todo/server/middleware/log"
	"go-todo/server/middleware/secure"
	"go-todo/server/validator"
	"net/http"
)

func NewEcho(cfg *config.Configuration) *echo.Echo {
	e := echo.New()
	logSkipper := func(c echo.Context) bool {
		return cfg.Server.SkipLogs
	}
	echoLogger := log.EchoLogger(logSkipper)

	bodyDumpSkipper := func(c echo.Context) bool {
		return cfg.Server.SkipBodyDump
	}
	bodyDumpLogger := log.BodyDumpLogger(bodyDumpSkipper)

	e.Validator = validator.NewEchoRequestValidator()

	e.Use(
		bodyDumpLogger,
		echoLogger,
		middleware.Recover(),
		secure.CORS(),
		secure.Headers(),
	)

	e.GET("/", healthCheck)

	return e
}

type response struct {
	Data string `json:"data"`
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, response{Data: "Go template at your service!🍲"})
}
