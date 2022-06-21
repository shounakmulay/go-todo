package routes

import (
	"net/http"

	errorutl "go-todo/internal/error"
	"go-todo/internal/json"
	"go-todo/server/controller"
	"go-todo/server/model/ctx"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
	"go-todo/server/validator"

	"github.com/labstack/echo/v4"
)

func User(g *echo.Group, controller controller.IUserController) {
	user := g.Group("/user")

	user.GET("/id/:id", func(c echo.Context) error {
		findUserReq := &reqmodel.UserID{}
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

	user.POST("", func(c echo.Context) error {
		reqUser := reqmodel.CreateUser{}
		bindValErr := validator.BindAndValidate(c, &reqUser)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		userID, err := controller.CreateUser(reqUser)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, reqUser.Username))
		}

		return json.Created(c, &resmodel.CreateUser{
			ID: userID,
		})
	})

	user.PUT("", func(c echo.Context) error {
		reqUser := reqmodel.UpdateUser{}
		bindValErr := validator.BindAndValidate(c, &reqUser)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		err := controller.UpdateUser(reqUser)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, reqUser.ID))
		}

		return c.NoContent(http.StatusOK)
	})

	user.DELETE("/id/:id", func(c echo.Context) error {
		userID := &reqmodel.UserID{}
		bindValErr := validator.BindAndValidate(c, userID)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		err := controller.DeleteUser(userID.ID)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, userID.ID))
		}

		return c.NoContent(http.StatusOK)
	})
}
