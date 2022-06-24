package redis

import (
	"go-todo/internal/log"
	"go-todo/server/config"

	"github.com/go-redis/redis/v8"
)

func Connect(cfg *config.Configuration) *redis.Client {
	log.Logger.Infof("Connecting to redis at %v", cfg.Redis.URL)

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.URL,
	})

	rdb.AddHook(redisLogger{})

	return rdb
}
