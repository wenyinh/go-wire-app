package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/wenyinh/go-wire-app/pkg/config"
)

type CacheClientInterface interface {
	RDB() *redis.Client
	Close() error
	Ping(context.Context) error
}

type CacheClient struct {
	Client *redis.Client
	config *config.RedisConfig
}

func NewCacheClient(cfg config.RedisConfig) (*CacheClient, func()) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return &CacheClient{
			Client: rdb,
			config: &cfg,
		}, func() {
			_ = rdb.Close()
		}
}

func (r *CacheClient) Ping(ctx context.Context) error {
	return r.Client.Ping(ctx).Err()
}

func (r *CacheClient) RDB() *redis.Client {
	return r.Client
}

func (r *CacheClient) Close() error {
	return r.RDB().Close()
}
