package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/badhu99/file-service/internal"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatalln("Check .env file: ", err)
	}

	server := configureServer()

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

	log.Printf("Starting %s server on port %s", internal.GetEnvString(server.Config.Env), srv.Addr)
	log.Fatal(srv.ListenAndServe())
}

func configureServer() *internal.Server {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatalln("Port could not be defined from .env file.", err)
	}

	environment := internal.Env(os.Getenv("ENVIRONMENT"))

	return &internal.Server{
		Config: internal.Config{
			Port: port,
			Env:  environment,
		},
	}
}
