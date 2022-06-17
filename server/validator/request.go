package validator

import (
	"github.com/labstack/echo/v4"
	"go-todo/server/model/resmodel"
)

func BindAndValidate(c echo.Context, model interface{}) *resmodel.ErrorResponse {
	bindErr := c.Bind(model)
	if bindErr != nil {
		return resmodel.UnprocessableEntity(bindErr)
	}
	valErr := c.Validate(model)
	if valErr != nil {
		return resmodel.BadRequest(valErr)
	}
	return nil
}
