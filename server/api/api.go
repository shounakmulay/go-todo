package api

import (
	"github.com/labstack/echo/v4"
	"go-todo/server/config"
	"gorm.io/gorm"
)

func Start(cfg *config.Configuration, db *gorm.DB) *echo.Echo {
	e := NewEcho()

	// TODO: Add Middlewares

	// TODO: Add routing

	return e
}
