package redis

import (
	"context"
	"github.com/go-redis/redis/v8"

	"time"
)

type RedisDB struct {
	db *redis.Client
}

func NewRedis(redis *redis.Client) RedisDB {
	return RedisDB{db: redis}
}

func (r *RedisDB) GetBytes(ctx context.Context, key string) ([]byte, error) {
	return r.db.Get(ctx, key).Bytes()
}

func (r *RedisDB) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	return r.db.Set(ctx, key, value, duration).Err()
}

func (r *RedisDB) Del(ctx context.Context, key string) error {
	return r.db.Del(ctx, key).Err()
}
