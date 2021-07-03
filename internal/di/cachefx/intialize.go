package cachefx

import (
	"github.com/go-redis/redis"

	"go-practice/pkg/cache"
)

func provideCache(redisCmd redis.UniversalClient) cache.Cache {
	return cache.NewRedisCache(redisCmd)
}
