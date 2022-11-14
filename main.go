package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/yoeritjuu/Computer-Store/pkg/config"
	"github.com/yoeritjuu/Computer-Store/pkg/handlers"
	"github.com/yoeritjuu/Computer-Store/pkg/repository"
)

func main() {
	godotenv.Load()
	config := config.LoadSqlConfig()
	sql, err := repository.ConnectMySQL(config)
	if err != nil {
		log.Fatalf("cant connect to mysql database, err is: %s", err.Error())
	}
	server := handlers.NewServer(config, sql)
	server.Run()
}
