package routes

import (
	errorutl "go-todo/internal/error"
	"go-todo/internal/json"
	"go-todo/server/controller"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
	"go-todo/server/validator"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func Auth(e *echo.Echo, userController controller.IUserController, jwtController controller.IJwtController) {
	auth := e.Group("/auth")

	auth.POST("/login", func(c echo.Context) error {
		reqLogin := &reqmodel.Login{}
		bindValErr := validator.BindAndValidate(c, reqLogin)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		user, err := userController.FindDBUserByUsername(reqLogin.Username)
		if err != nil {
			return json.Error(c, errorutl.GormToResErr(err, reqLogin.Username))
		}

		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqLogin.Password))
		if passErr != nil {
			incorrectPassError := errors.WithMessage(passErr, "Incorrect Password")
			return json.Error(c, resmodel.Unauthorized(incorrectPassError))
		}

		jwtTokens, jwtError := jwtController.GenerateTokens(user)
		if jwtError != nil {
			return json.Error(c, resmodel.InternalServerError(jwtError))
		}

		user.Token = jwtTokens.RefreshToken
		updateErr := userController.UpdateUserToken(&user)
		if updateErr != nil {
			return json.Error(c, resmodel.InternalServerError(updateErr))
		}

		return json.Success(c, jwtTokens)
	})
}
