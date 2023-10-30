package handler

import (
	"log"
	"net/http"
)

func (handlerData *HandlerData) ValidateAccessToken(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(string)
	if !ok {
		log.Println(userId)
		log.Println(ok)
		http.Error(w, "Could not get userId from bearer token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(userId))
}
