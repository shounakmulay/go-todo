package jwt

import (
	"errors"
	"fmt"
	"go-todo/server/config"
	"go-todo/server/controller"
	"go-todo/server/model/claims"
	"regexp"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	adminPathsRegex = "(\\/api\\/user.*)|(\\/api\\/role.*)"

	adminRoleId = 2
)

func JWT(cfg *config.JWT, controller controller.IUserController) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: cfg.Secret,
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {

			// Parse token
			token, err := jwt.ParseWithClaims(auth, &claims.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
				if token.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
				}
				return []byte(cfg.Secret), nil
			})
			if err != nil {
				return nil, err
			}

			// Check token validity and extract jwtClaims
			if jwtClaims, ok := token.Claims.(*claims.JwtClaims); ok && token.Valid {
				roleID := jwtClaims.Role
				// username := claims.Username
				path := c.Path()

				// Check if request requires admin
				isAdminPath, err := regexp.MatchString(adminPathsRegex, path)
				if err != nil {
					// Entering this block means the regex is invalid.
					// This should never happen, thus we panic instead of returning err.
					panic(fmt.Sprintf("Invalid regex: %v", err))
				}

				if isAdminPath && roleID != adminRoleId {
					return nil, errors.New("unauthorized! Only admins are authorized to make this request")
				}

				// // Check if the user exists in the database
				// _, dbErr := controller.FindDBUserByUsername(username)
				// if dbErr != nil {
				// 	return nil, errors.New("unauthorized! User not found")
				// }

				return token, nil
			} else {
				return nil, errors.New("could not validate token")
			}
		},
	})
}
