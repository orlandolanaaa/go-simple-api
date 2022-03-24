package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/go-redis/redismock/v8"
	"reflect"
	"testing"
	"time"
)

var (
	client *redis.Client
)

var (
	key = "key"
	val = "val"
)

func TestNewRedis(t *testing.T) {
	db, _ := redismock.NewClientMock()
	tests := []struct {
		name string
		want RedisDB
	}{
		{name: "Initiate-Redis", want: NewRedis(db)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRedis(db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisDB_Del(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
	}
	db, mock := redismock.NewClientMock()
	newsID := "value"

	ctx := context.Background()
	key := fmt.Sprintf("news_redis_cache_%s", newsID)
	mock.ExpectDel(key).SetErr(nil)
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Delete-Redis", args: args{
			ctx: ctx,
			key: key,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(db)
			_ = r.Del(ctx, tt.args.key)
		})
	}
}

func TestRedisDB_GetBytes(t *testing.T) {
	db, mock := redismock.NewClientMock()
	newsID := "value"

	ctx := context.Background()
	key := fmt.Sprintf("news_redis_cache_%s", newsID)
	mock.ExpectGet(key).SetErr(nil)

	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{name: "Get-Redis", args: args{
			ctx: context.Background(),
			key: key,
		}, want: []byte(key), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(db)
			_, _ = r.GetBytes(ctx, tt.args.key)

		})
	}
}

func TestRedisDB_Set(t *testing.T) {
	type args struct {
		ctx      context.Context
		key      string
		value    interface{}
		duration time.Duration
	}
	db, mock := redismock.NewClientMock()
	newsID := "value"
	key := fmt.Sprintf("news_redis_cache_%s", newsID)
	mock.ExpectSet(key, newsID, 1*time.Second).SetVal(newsID)

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Set-Redis", args: args{
			ctx:      context.Background(),
			key:      key,
			value:    newsID,
			duration: 1 * time.Second,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRedis(db)
			if err := r.Set(tt.args.ctx, tt.args.key, tt.args.value, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
