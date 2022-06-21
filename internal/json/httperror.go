package json

import (
	"go-todo/server/model/resmodel"

	"github.com/labstack/echo/v4"
)

func Error(c echo.Context, err *resmodel.ErrorResponse) error {
	return c.JSON(err.Status, err)
}
