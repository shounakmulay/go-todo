package ratelimit

import (
	"context"
	"go-todo/server/config"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	rateLimitBaseKey = "rate-limit-"
)

type AuthRateLimitStore struct {
	rdb   *redis.Client
	ttl   time.Duration
	limit int
	mutex sync.Mutex
}

func NewAuthRateLimitStore(rdb *redis.Client, cfg *config.Redis) *AuthRateLimitStore {
	return &AuthRateLimitStore{
		rdb:   rdb,
		ttl:   time.Duration(cfg.AuthRateLimitDurationSeconds) * time.Second,
		limit: cfg.AuthRateLimitCount,
		mutex: sync.Mutex{},
	}
}

// Stores for the rate limiter have to implement the Allow method
func (rs *AuthRateLimitStore) Allow(identifier string) (bool, error) {
	rs.mutex.Lock()
	defer rs.mutex.Unlock()

	ctx := context.Background()
	key := rateLimitBaseKey + identifier
	visits, err := rs.rdb.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if visits > int64(rs.limit) {
		return false, nil
	}

	if visits == 1 {
		setErr := rs.rdb.Set(ctx, key, 1, rs.ttl).Err()
		if setErr != nil {
			return false, setErr
		}
	}

	return true, nil
}
