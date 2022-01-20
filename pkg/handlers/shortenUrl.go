package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/rs/xid"
)

type urlPayload struct {
	Url string `json:"url"`
}

func (h *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	log.Println("requesting at path: shorten-url")
	var payload urlPayload

	json.NewDecoder(r.Body).Decode(&payload)
	jsonResponseEncoder := json.NewEncoder(w)

	if payload.Url == "" {
		w.WriteHeader(http.StatusBadRequest)
		jsonResponseEncoder.Encode(map[string]string{
			"message": "no url supplied",
		})

		return
	}

	key := xid.New().String()

	log.Printf("generated key: %v", key)

	err := h.redis.Set(key, payload.Url, 24*time.Hour)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonResponseEncoder.Encode(map[string]string{
		"shortened_key": key,
	})
}
