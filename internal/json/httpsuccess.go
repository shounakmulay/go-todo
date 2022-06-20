package json

import (
	"github.com/labstack/echo/v4"
	"go-todo/server/model/resmodel"
	"net/http"
)

func Success(c echo.Context, data any) error {
	response := resmodel.Response{
		Status: http.StatusOK,
		Body:   data,
	}
	return c.JSON(http.StatusOK, response)
}

func Created(c echo.Context, data any) error {
	response := resmodel.Response{
		Status: http.StatusCreated,
		Body:   data,
	}
	return c.JSON(http.StatusCreated, response)
}
