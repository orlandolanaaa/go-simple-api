package main

import (
	"be_entry_task/cmd/server"
	"be_entry_task/internal/mysql"
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
	mysql.InitCon()
	srv := server.Get().WithAddr(os.Getenv("PORT")).
		WithRouter(server.ListRoute())

	if err := srv.Start(); err != nil {
		fmt.Printf(err.Error())
	}
}
