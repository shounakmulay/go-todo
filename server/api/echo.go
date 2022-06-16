package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-todo/server/middleware/secure"
	"net/http"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover(), secure.CORS(), secure.Headers())
	e.GET("/", healthCheck)
	return e
}

type response struct {
	Data string `json:"data"`
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, response{Data: "Go template at your service!🍲"})
}
