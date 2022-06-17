package json

import (
	"github.com/labstack/echo/v4"
	"go-todo/server/model/resmodel"
)

func Error(c echo.Context, err *resmodel.ErrorResponse) error {
	return c.JSON(err.Status, err)
}
