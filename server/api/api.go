package api

import (
	"go-todo/server/cache"
	"go-todo/server/config"
	"go-todo/server/controller"
	"go-todo/server/daos"
	"go-todo/server/middleware/jwt"
	"go-todo/server/middleware/ratelimit"
	"go-todo/server/routes"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Start(cfg *config.Configuration, db *gorm.DB, redis *redis.Client) (*echo.Echo, error) {
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

	// Cache
	userCache := cache.NewUserCache(redis)
	todoCache := cache.NewTodoCache(redis)

	// Controllers
	roleController := controller.NewRoleController(roleDao)
	userController := controller.NewUserController(userDao, userCache)
	todoController := controller.NewTodoController(todoDao, todoCache)

	// Auth Middlewares
	authRateLimiter := ratelimit.AuthRateLimit(redis, cfg.Redis)

	// Auth Routes
	routes.Auth(e, userController, jwtController, authRateLimiter)

	// API middlewares
	jwtMiddleware := jwt.JWT(cfg.JWT, userController)

	// API Routes
	api := e.Group("/api", jwtMiddleware)

	routes.Role(api, roleController)
	routes.User(api, userController)
	routes.Todo(api, todoController)

	return e, nil
}
