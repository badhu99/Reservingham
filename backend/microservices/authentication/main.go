package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	server "github.com/badhu99/authentication/internal"
	"github.com/gorilla/handlers"
)

func main() {

	server := &server.Server{
		Config: server.Config{
			Port: 8080,
			Env:  server.Dev,
		},
		Database: server.DatabaseConnect(),
	}

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", server.Config.Port),
		Handler:      handlers.CORS(origins, header, methods)(server.AuthRoutes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting %s server on port %s", server.Config.Env, srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
