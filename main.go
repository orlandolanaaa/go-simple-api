package main

import (
	"be_entry_task/cmd/server"
	"be_entry_task/internal/mysql"
	"be_entry_task/internal/redis"
	_ "flag"
	"fmt"
	_ "github.com/docker/docker/daemon/logger"
	_ "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "go.uber.org/zap"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect mysql
	db, err := mysql.Conn()
	defer db.Close()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redis := redis.Initiate(os.Getenv("REDIS_DB_HOST") + ":" + os.Getenv("REDIS_DB_PORT"))
	srv := server.Get(db, redis).WithAddr(os.Getenv("PORT")).
		WithRouter()

	if err := srv.Start(); err != nil {
		fmt.Printf(err.Error())
	}

}
