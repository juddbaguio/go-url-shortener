package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) RedirectWithKey(w http.ResponseWriter, r *http.Request) {
	log.Println("requesting at path: redirect-with-key")
	vars := mux.Vars(r)

	url, err := h.redis.Get(vars["key"])

	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
