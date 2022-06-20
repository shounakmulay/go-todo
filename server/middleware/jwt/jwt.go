package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-todo/server/config"
	"go-todo/server/controller"
	"go-todo/server/model"
	"regexp"
)

var (
	adminPathsRegex = "(\\/api\\/user.*)"

	adminRoleId = 2
)

func JWT(cfg *config.JWT, controller controller.IUserController) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: cfg.Secret,
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			// TODO: Check if route requires auth

			claims := &model.JwtCustomClaims{}
			token, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
				if token.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
				}
				return []byte(cfg.Secret), nil
			})
			if err != nil {
				return nil, err
			}

			if claims, ok := token.Claims.(*model.JwtCustomClaims); ok && token.Valid {
				roleID := claims.Role
				//username := claims.Username
				path := c.Path()

				isAdminPath, err := regexp.MatchString(adminPathsRegex, path)
				if err != nil {
					// Entering this block means the regex is invalid.
					// This should never happen, thus we panic instead of returning err.
					panic(fmt.Sprintf("Invalid regex: %v", err))
				}

				if isAdminPath && roleID != adminRoleId {
					return nil, errors.New("Unauthorized! \n Only admins are authorized to make this request.")
				}

				// TODO: Check if user exists

				return token, nil
			} else {
				return nil, errors.New("could not validate token")
			}
		},
	})
}
