package infra

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/juddbaguio/url-shortener/pkg/config"
)

var ctx = context.Background()

type Redis struct {
	rdb *redis.Client
}

type RedisService interface {
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
}

func NewRedisClient(cfg config.Cfg) (RedisService, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("redis:%v", cfg.RedisPort),
		Password:    "best-url-shortener",
		DB:          0,
		DialTimeout: 15 * time.Second,
	})

	_, err := rdb.Ping(context.Background()).Result()

	return &Redis{
		rdb: rdb,
	}, err
}

func (r *Redis) Set(key string, value interface{}, exp time.Duration) error {
	return r.rdb.Set(ctx, key, value, exp).Err()
}

func (r *Redis) Get(key string) (string, error) {
	return r.rdb.Get(ctx, key).Result()
}
