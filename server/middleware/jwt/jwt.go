package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-todo/server/config"
	"go-todo/server/controller"
	"go-todo/server/model"
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
				return cfg.Secret, nil
			})
			if err != nil {
				return nil, err
			}

			if claims, ok := token.Claims.(*model.JwtCustomClaims); ok && token.Valid {
				// TODO: Check if claims valid for path. Check role valid for path.
				// TODO: Check if user exists
				panic(claims)
			} else {
				// TODO: Handle invalid token
				panic("")
			}
		},
	})
}
