package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

var Client *redis.Client

func Initiate(addr string) *redis.Client {
	Client = redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("error init redis client: #{err}\n", err)
	}
	fmt.Println("Redis started !")

	return Client
}
