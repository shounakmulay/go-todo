package ratelimit

import (
	"errors"
	"fmt"
	"go-todo/server/config"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthRateLimit(rdb *redis.Client, cfg *config.Redis) echo.MiddlewareFunc {
	rateLimiterConfig := middleware.RateLimiterConfig{
		Skipper:             func(echo.Context) bool { return cfg.DisableRateLimit },
		IdentifierExtractor: identifierExtractor,
		Store:               NewAuthRateLimitStore(rdb, cfg),
		ErrorHandler:        errorHandler,
		DenyHandler:         denyHandler,
	}

	return middleware.RateLimiterWithConfig(rateLimiterConfig)
}

func identifierExtractor(c echo.Context) (string, error) {
	userIP := c.RealIP()
	if userIP == "" {
		return "", errors.New("unable to read client IP")
	}
	path := c.Path()

	return path + "-" + userIP, nil
}

func errorHandler(c echo.Context, err error) error {
	return &echo.HTTPError{
		Code:     http.StatusForbidden,
		Message:  "error processing request",
		Internal: err,
	}
}

func denyHandler(c echo.Context, identifier string, err error) error {
	return &echo.HTTPError{
		Code:     http.StatusTooManyRequests,
		Message:  fmt.Sprintf("rate limit for %s exceeded", c.Path()),
		Internal: err,
	}
}
