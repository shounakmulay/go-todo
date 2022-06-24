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

type UserCache struct {
	rdb          *redis.Client
	cacheBaseKey string
	ttl          time.Duration
}

func NewUserCache(rdb *redis.Client) *UserCache {
	return &UserCache{
		rdb:          rdb,
		cacheBaseKey: "usercache_",
		ttl:          time.Minute * 5,
	}
}

func (c *UserCache) GetUserByUsername(username string) (*resmodel.User, error) {
	cacheKey := c.cacheBaseKey + username
	return c.getUser(cacheKey)
}

func (c *UserCache) GetUserByID(id int) (*resmodel.User, error) {
	cacheKey := c.cacheBaseKey + strconv.Itoa(id)
	return c.getUser(cacheKey)
}

func (c *UserCache) SetUserForUsername(username string, user resmodel.User) error {
	cacheKey := c.cacheBaseKey + username
	return c.setUser(cacheKey, user)
}

func (c *UserCache) SetUserForID(id int, user resmodel.User) error {
	cacheKey := c.cacheBaseKey + strconv.Itoa(id)
	return c.setUser(cacheKey, user)
}

func (c *UserCache) Invalidate(id int, username string) error {
	ctx := context.Background()
	_, err := c.rdb.Del(ctx, c.cacheBaseKey+strconv.Itoa(id), c.cacheBaseKey+username).Result()
	return err
}

func (c *UserCache) getUser(key string) (*resmodel.User, error) {
	ctx := context.Background()
	val, err := c.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	resUser := &resmodel.User{}
	jsonErr := json.Unmarshal([]byte(val), resUser)
	if jsonErr != nil {
		// Could not unmarshal json. Deleting key.
		errutl.Log(c.rdb.Del(ctx, key).Err())
		return nil, jsonErr
	}

	return resUser, nil
}

func (c *UserCache) setUser(key string, user resmodel.User) error {
	ctx := context.Background()

	jsonUser, jsonErr := json.Marshal(user)
	if jsonErr != nil {
		return jsonErr
	}

	_, err := c.rdb.Set(ctx, key, string(jsonUser), c.ttl).Result()
	return err
}
