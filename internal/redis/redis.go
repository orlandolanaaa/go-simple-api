package redis

import (
	"context"
	"time"
)

type RedisDB struct {
}

func NewRedis() RedisDB {
	return RedisDB{}
}

func (r *RedisDB) GetBytes(ctx context.Context, key string) ([]byte, error) {
	return Client.Get(ctx, key).Bytes()
}

func (r *RedisDB) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	return Client.Set(ctx, key, value, duration).Err()
}

func (r *RedisDB) Del(ctx context.Context, key string) error {
	return Client.Del(ctx, key).Err()
}
