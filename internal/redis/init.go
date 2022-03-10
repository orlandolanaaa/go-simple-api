package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

var Client *redis.Client

func Initiate() {
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf(os.Getenv("REDIS_DB_HOST") + ":" + os.Getenv("REDIS_DB_PORT")),
		Password: os.Getenv("REDIS_DB_PASSWORD"),
		DB:       0,
	})

	_, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln("error init redis client: #{err}\n", err)
	}
	fmt.Println("Redis started !")
}
