package server

import (
	"log"
	"os"

	"github.com/badhu99/api-gateway/internal/handler"
	"github.com/badhu99/api-gateway/internal/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Server struct {
	Config Config
}

type Config struct {
	Port int
	Env  Env
}

type Env string

const (
	Dev  = "development"
	Test = "testing"
	Prod = "production"
)

func (server *Server) AuthRoutes() *mux.Router {

	router := mux.NewRouter()

	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatalln("Check .env file: ", err)
	}

	urlAuth := os.Getenv("URL_AUTH")

	h := handler.HandlerData{
		UrlAuth:       urlAuth,
		UrlManagement: "http://localhost:8082",
	}

	router = router.PathPrefix("/api").Subrouter()
	router.Use(middleware.Log)

	routerCompany := router.PathPrefix("/auth").Subrouter()
	routerCompany.HandleFunc("", h.SignIn).Methods("POST")

	return router
}
