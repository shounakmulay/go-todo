package api

import (
	"github.com/labstack/echo/v4"
	"go-todo/server/config"
	"go-todo/server/controller"
	"go-todo/server/daos"
	"go-todo/server/routes"
	"gorm.io/gorm"
)

func Start(cfg *config.Configuration, db *gorm.DB) (*echo.Echo, error) {
	e := NewEcho(cfg)

	// TODO: Add Middlewares

	// TODO: Add routing

	// Setup Dependencies
	roleDao := daos.NewDao(db)
	roleController := controller.NewRoleController(roleDao)

	api := e.Group("/api")
	routes.Role(api, roleController)

	return e, nil
}
