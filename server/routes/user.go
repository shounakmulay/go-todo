package routes

import (
	"github.com/labstack/echo/v4"
	errorutl "go-todo/internal/error"
	"go-todo/internal/json"
	"go-todo/server/controller"
	"go-todo/server/model/reqmodel"
	"go-todo/server/validator"
)

func User(g *echo.Group, controller controller.IUserController) {
	user := g.Group("/user")

	user.GET("/id/:id", func(c echo.Context) error {
		findUserReq := &reqmodel.FindUserByID{}
		bindValErr := validator.BindAndValidate(c, findUserReq)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		user, err := controller.FindUser(findUserReq.ID)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, findUserReq.ID))
		}

		return json.Success(c, user)
	})

	user.GET("/username/:username", func(c echo.Context) error {
		findUserReq := &reqmodel.FindUserByUsername{}
		bindValErr := validator.BindAndValidate(c, findUserReq)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		user, err := controller.FindUserByUsername(findUserReq.Username)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, findUserReq.Username))
		}

		return json.Success(c, user)
	})
}
