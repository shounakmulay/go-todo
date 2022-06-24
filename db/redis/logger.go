package redis

import (
	"context"
	"go-todo/internal/log"

	"github.com/go-redis/redis/v8"
)

type redisLogger struct{}

func (r redisLogger) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (r redisLogger) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	log.Logger.Infof("REDIS: %s", cmd.String())
	return nil
}

func (r redisLogger) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (r redisLogger) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	var cmdsStr string
	for _, cmd := range cmds {
		cmdsStr += cmd.String() + "\n"
	}
	log.Logger.Infof("REDIS: %v", cmdsStr)
	return nil
}
