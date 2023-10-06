package router

import (
	"log"
	"net/http"

	"github.com/badhu99/authentication/internal/handler"
	"github.com/gorilla/mux"
)

var r *mux.Router

func RegisterRoutes() {
	log.Println("Listening on port 8080")

	r = mux.NewRouter()

	r.HandleFunc("/auth/signin", handler.SignIn).Methods("GET")

	http.Handle("/", r)

	http.ListenAndServe("0.0.0.0:8080", r)
	log.Println("Listening on port 8080")
}
