package routes

import (
	"github.com/labstack/echo/v4"
	errorutl "go-todo/internal/error"
	"go-todo/internal/json"
	"go-todo/server/controller"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
	"go-todo/server/validator"
)

func Role(g *echo.Group, controller controller.IRoleController) {
	role := g.Group("/role")

	role.GET("/:id", func(c echo.Context) error {
		role := &reqmodel.Role{}
		bindValErr := validator.BindAndValidate(c, role)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		dbRole, err := controller.GetRole(role.ID)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, role.ID))
		}

		resRole := resmodel.Role{
			ID:          dbRole.ID,
			Name:        dbRole.Name,
			AccessLevel: dbRole.AccessLevel,
		}
		return json.Success(c, resRole)
	})
}
