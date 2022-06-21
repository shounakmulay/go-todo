package ctx

import (
	"go-todo/server/config"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	RoleID   int    `json:"roleId"`
}

func GetUserFromContext(c echo.Context) User {
	user := c.Get(config.JwtContextKey).(User)
	return user
}
