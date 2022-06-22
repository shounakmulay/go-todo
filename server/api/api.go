package api

import (
	"go-todo/server/config"
	"go-todo/server/controller"
	"go-todo/server/daos"
	"go-todo/server/middleware/jwt"
	"go-todo/server/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Start(cfg *config.Configuration, db *gorm.DB) (*echo.Echo, error) {
	e := NewEcho(cfg)

	// Setup Dependencies
	jwtController, err := controller.NewJwtController(cfg.JWT)
	if err != nil {
		return nil, err
	}

	// Daos
	roleDao := daos.NewRoleDao(db)
	userDao := daos.NewUserDao(db)
	todoDao := daos.NewTodoDao(db)

	// Controllers
	roleController := controller.NewRoleController(roleDao)
	userController := controller.NewUserController(userDao)
	todoController := controller.NewTodoController(todoDao)

	// Routes without JWT
	routes.Auth(e, userController, jwtController)

	jwtMiddleware := jwt.JWT(cfg.JWT, userController)
	api := e.Group("/api", jwtMiddleware)

	// Routes with JWT
	routes.Role(api, roleController)
	routes.User(api, userController)
	routes.Todo(api, todoController)

	return e, nil
}
