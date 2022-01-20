package config

import (
	"fmt"
	"os"
)

type Cfg struct {
	Port          string
	RedisPort     string
	RedisPassword string
}

func Get() (Cfg, error) {
	port := os.Getenv("PORT")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	if port == "" {
		return Cfg{}, fmt.Errorf("port is missing")
	}

	if redisPort == "" {
		return Cfg{}, fmt.Errorf("redis port is missing")
	}

	if redisPassword == "" {
		return Cfg{}, fmt.Errorf("redis password is missing")
	}

	return Cfg{
		Port:          port,
		RedisPort:     redisPort,
		RedisPassword: redisPassword,
	}, nil
}
