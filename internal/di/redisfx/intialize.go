package redisfx

import (
	"context"

	"github.com/go-redis/redis"
	"go.uber.org/fx"

	"github.com/spf13/viper"
)

func provideRedisCmd(lifecycle fx.Lifecycle) redis.UniversalClient {
	opts := &redis.Options{
		Addr:     viper.GetString("REDIS_URL"),
		Password: viper.GetString("REDIS_PASSWORD"),
	}
	client := redis.NewClient(opts)
	lifecycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})
	return client
}
