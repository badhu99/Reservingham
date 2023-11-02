package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	server "github.com/badhu99/management-service/internal"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatalln("Check .env file: ", err)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalln("Port could not be defined from .env file.", err)
	}

	server := &server.Server{
		Config: server.Config{
			Port: port,
			Env:  server.Dev,
		},
		Database: server.DatabaseConnect(),
	}

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", server.Config.Port),
		Handler:      handlers.CORS(origins, header, methods)(server.Routes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("Starting %s server on port %s", server.Config.Env, srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
