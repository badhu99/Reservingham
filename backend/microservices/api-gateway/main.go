package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/badhu99/api-gateway/docs"
	server "github.com/badhu99/api-gateway/internal"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample Reservingham server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type 'Bearer TOKEN' to correctly set the API Key
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
