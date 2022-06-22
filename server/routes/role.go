package routes

import (
	errorutl "go-todo/internal/error"
	"go-todo/internal/json"
	"go-todo/server/controller"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
	"go-todo/server/validator"

	"github.com/labstack/echo/v4"
)

func Role(g *echo.Group, controller controller.IRoleController) {
	role := g.Group("/role")

	role.POST("/", func(c echo.Context) error {
		role := &reqmodel.CreateRole{}
		bindValErr := validator.BindAndValidateWith(c, role, validator.BindBody)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		dbRoleID, err := controller.CreateRole(*role)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, dbRoleID))
		}

		return json.Success(c, resmodel.RoleID{
			ID: dbRoleID,
		})
	})

	role.GET("/:id", func(c echo.Context) error {
		role := &reqmodel.FindRole{}
		bindValErr := validator.BindAndValidateWith(c, role, validator.BindPath)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		resRole, err := controller.FindRoleByID(role.ID)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, role.ID))
		}

		return json.Success(c, resRole)
	})
}
