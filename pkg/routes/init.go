package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juddbaguio/url-shortener/pkg/handlers"
	"github.com/juddbaguio/url-shortener/pkg/infra"
)

func SetUpRoutes(mux *mux.Router, redis infra.RedisService) {
	handlers := handlers.InitHandlers(redis)

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(map[string]string{
			"message": "Welcome!",
		})
	}).Methods("GET")
	mux.HandleFunc("/", handlers.ShortenUrl).Methods("POST")
	mux.HandleFunc("/{key}", handlers.RedirectWithKey).Methods("GET")
}
