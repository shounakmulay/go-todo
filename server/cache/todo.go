package cache

import (
	"context"
	"encoding/json"
	errutl "go-todo/internal/error"
	"go-todo/server/model/resmodel"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type TodoCache struct {
	rdb          *redis.Client
	cacheBaseKey string
	ttl          time.Duration
}

func NewTodoCache(rdb *redis.Client) *TodoCache {
	return &TodoCache{
		rdb:          rdb,
		cacheBaseKey: "todoscache_",
		ttl:          time.Minute * 5,
	}
}

func (c TodoCache) GetAllTodos(userID int) ([]resmodel.Todo, error) {
	key := c.cacheBaseKey + strconv.Itoa(userID)
	ctx := context.Background()
	result, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	todos := []resmodel.Todo{}
	jsonErr := json.Unmarshal([]byte(result), &todos)
	if jsonErr != nil {
		// Could not unmarshal json. Deleting key.
		errutl.Log(c.rdb.Del(ctx, key).Err())
		return nil, jsonErr
	}

	return todos, nil
}

func (c TodoCache) SetAllTodos(userID int, todos []resmodel.Todo) error {
	key := c.cacheBaseKey + strconv.Itoa(userID)

	jsonBytes, jsonErr := json.Marshal(todos)
	if jsonErr != nil {
		return jsonErr
	}

	ctx := context.Background()

	return c.rdb.Set(ctx, key, string(jsonBytes), c.ttl).Err()
}

func (c TodoCache) Invalidate(userID int) error {
	ctx := context.Background()
	key := c.cacheBaseKey + strconv.Itoa(userID)
	return c.rdb.Del(ctx, key).Err()
}
