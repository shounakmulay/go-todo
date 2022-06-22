package routes

import (
	"fmt"

	errorutl "go-todo/internal/error"
	"go-todo/internal/json"
	"go-todo/server/controller"
	"go-todo/server/model/claims"
	"go-todo/server/model/dbmodel"
	"go-todo/server/model/reqmodel"
	"go-todo/server/model/resmodel"
	"go-todo/server/validator"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func Auth(e *echo.Echo, userController controller.IUserController, jwtController controller.IJwtController) {
	auth := e.Group("/auth")

	auth.POST("/login", func(c echo.Context) error {
		reqLogin := &reqmodel.Login{}
		bindValErr := validator.BindAndValidateWith(c, reqLogin, validator.BindBody)
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

		jwtTokens, jwtError := generateTokensAndUpdateUser(c, user, userController, jwtController)
		if jwtError != nil {
			return json.Error(c, jwtError)
		}

		return json.Success(c, jwtTokens)
	})

	auth.POST("/refresh", func(c echo.Context) error {
		reqRefresh := &reqmodel.Refresh{}
		bindValErr := validator.BindAndValidateWith(c, reqRefresh, validator.BindBody)
		if bindValErr != nil {
			return json.Error(c, bindValErr)
		}

		refreshToken, err := jwt.ParseWithClaims(
			reqRefresh.RefreshToken,
			&claims.JwtRefreshClaims{},
			func(token *jwt.Token) (interface{}, error) {
				if token.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
				}
				return jwtController.GetRefreshSecret(), nil
			})

		if err != nil {
			return json.Error(c, resmodel.Unauthorized(err))
		}

		if refClaims, ok := refreshToken.Claims.(*claims.JwtRefreshClaims); ok && refreshToken.Valid {
			user, err := userController.FindDBUserByUsername(refClaims.Username)
			if err != nil {
				return json.Error(c, errorutl.GormToResErr(err, refClaims.Username))
			}

			if refClaims.Subject != "refresh" {
				return json.Error(c, resmodel.Unauthorized(errors.New("Invalid refresh token")))
			}

			if user.ID != refClaims.ID {
				return json.Error(c, resmodel.Unauthorized(errors.New("Invalid refresh token for user")))
			}

			if user.Token != reqRefresh.RefreshToken {
				return json.Error(c, resmodel.Unauthorized(errors.New("Invalid refresh token")))
			}

			jwtTokens, jwtError := generateTokensAndUpdateUser(c, user, userController, jwtController)
			if jwtError != nil {
				return json.Error(c, jwtError)
			}

			return json.Success(c, jwtTokens)
		}

		return json.Error(c, resmodel.Unauthorized(errors.New("Invalid refresh token")))
	})
}

func generateTokensAndUpdateUser(
	c echo.Context, user dbmodel.User,
	userController controller.IUserController,
	jwtController controller.IJwtController,
) (resmodel.JwtTokens, *resmodel.ErrorResponse) {
	jwtTokens, jwtError := jwtController.GenerateTokens(user)
	if jwtError != nil {
		return jwtTokens, resmodel.InternalServerError(jwtError)
	}

	user.Token = jwtTokens.RefreshToken
	updateErr := userController.UpdateUserToken(&user)

	if updateErr != nil {
		return jwtTokens, resmodel.InternalServerError(updateErr)
	}

	return jwtTokens, nil
}
