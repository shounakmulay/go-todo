package api

import (
	"github.com/labstack/echo/v4"
	"go-todo/server/config"
	"go-todo/server/controller"
	"go-todo/server/daos"
	"go-todo/server/middleware/jwt"
	"go-todo/server/routes"
	"gorm.io/gorm"
)

func Start(cfg *config.Configuration, db *gorm.DB) (*echo.Echo, error) {
	e := NewEcho(cfg)

	// TODO: Add Middlewares

	// TODO: Add routing

	// Setup Dependencies
	jwtController, err := controller.NewJwtController(cfg.JWT)
	if err != nil {
		return nil, err
	}

	roleDao := daos.NewDao(db)
	userDao := daos.NewUserDao(db)

	roleController := controller.NewRoleController(roleDao)
	userController := controller.NewUserController(userDao)

	routes.Auth(e, userController, jwtController)

	jwtMiddleware := jwt.JWT(cfg.JWT, userController)
	api := e.Group("/api", jwtMiddleware)

	routes.Role(api, roleController)
	routes.User(api, userController)

	return e, nil
}
