package cache

import (
	"github.com/go-redis/redis"
)

type redisCache struct {
	redisCmd redis.UniversalClient
}

func (r *redisCache) Get(key string, target interface{}) error {
	panic("implement me")
}

func (r *redisCache) Set(key string, data interface{}) error {
	panic("implement me")
}

func NewRedisCache(redisCmd redis.UniversalClient) Cache {
	return &redisCache{redisCmd: redisCmd}
}
