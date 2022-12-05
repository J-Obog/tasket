package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (this *RedisCache) Get(key string) (*string, error) {
	val, err := this.client.Get(context.Background(), key).Result()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}

		return nil, err
	}

	res := new(string)
	*res = val

	return res, nil
}

func (this *RedisCache) Set(key string, value string, expr int64) error {
	exprTime := time.Duration(expr) * time.Second
	return this.client.SetEx(context.Background(), key, value, exprTime).Err()
}
