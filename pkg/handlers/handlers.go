package handlers

import "github.com/juddbaguio/url-shortener/pkg/infra"

type Handler struct {
	redis infra.RedisService
}

func InitHandlers(redis infra.RedisService) *Handler {
	return &Handler{
		redis: redis,
	}
}
