package handler

import (
	"log"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	log.Println("Pinged")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("dejla"))
}
