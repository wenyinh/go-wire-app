package redis

import (
	"context"
	goRedis "github.com/redis/go-redis/v9"
	"github.com/wenyinh/go-wire-app/pkg/config"
)

type CacheClient struct {
	Client *goRedis.Client
}

func NewCacheClient(cfg config.RedisConfig) *CacheClient {
	rdb := goRedis.NewClient(&goRedis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return &CacheClient{Client: rdb}
}

func (r *CacheClient) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}
