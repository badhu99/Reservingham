package main

import (
	"log"
	"net/http"

	"github.com/badhu99/api-gateway/handlers"
	"github.com/gorilla/mux"
)

func main() {

	log.Println("Listening on port 8081")

	r := mux.NewRouter()

	r.HandleFunc("/api/auth/signin", handlers.SingIn)

	http.Handle("/", r)

	http.ListenAndServe("0.0.0.0:8081", r)
}
