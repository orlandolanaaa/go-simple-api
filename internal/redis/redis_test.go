package redis

import (
	"context"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
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
	tests := []struct {
		name string
		want RedisDB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := NewRedis(); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewRedis() = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestRedisDB_Del(t *testing.T) {
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisDB{}
			if err := r.Del(tt.args.ctx, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Del() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisDB_Get(t *testing.T) {
	mock := redismock.NewNiceMock(client)
	mock.On("Get", key).Return(redis.NewStringResult(val, nil))

	r := NewRedis(mock)
	res, err := r.Get(context.Background(), key)
	assert.NoError(t, err)
	assert.Equal(t, val, res)
}

func TestRedisDB_GetBytes(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisDB{}
			got, err := r.GetBytes(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBytes() got = %v, want %v", got, tt.want)
			}
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
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RedisDB{}
			if err := r.Set(tt.args.ctx, tt.args.key, tt.args.value, tt.args.duration); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
