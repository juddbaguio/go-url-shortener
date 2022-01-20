package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/juddbaguio/url-shortener/pkg/api"
	"github.com/juddbaguio/url-shortener/pkg/config"
	"github.com/juddbaguio/url-shortener/pkg/infra"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Printf("error loading environment variables")
		os.Exit(1)
	}

	config, err := config.Get()

	if err != nil {
		log.Printf("environment variable error: %v", err.Error())
		os.Exit(1)
	}

	redisSrv, err := infra.NewRedisClient(config)
	if err != nil {
		log.Printf("redis error: %v", err.Error())
		os.Exit(1)
	}

	server := api.NewServer(redisSrv)

	if err := server.Start(config); err != nil {
		log.Printf("server error: %v\n", err.Error())
		os.Exit(1)
	}
}
