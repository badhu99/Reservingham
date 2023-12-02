package internal

import (
	"log"
	"os"

	"github.com/badhu99/file-service/internal/handler"
	"github.com/badhu99/file-service/internal/middleware"
	"github.com/badhu99/file-service/internal/services"
	"github.com/gorilla/mux"
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
	Dev  = "DEV"
	Test = "TEST"
	Prod = "PROD"
)

func GetEnvString(env Env) string {
	switch env {
	case Dev:
		return "development"
	case Test:
		return "testing"
	case Prod:
		return "production"
	default:
		log.Fatalf("Invalid environment")
		return ""
	}
}

func (server *Server) Routes() *mux.Router {
	router := mux.NewRouter()

	// File storage init props
	filePath := os.Getenv("PORT")

	h := handler.HandlerData{
		FileStorage: services.NewLocalFileStorage(filePath),
	}

	fileRouter := router.PathPrefix("/file").Subrouter()
	fileRouter.HandleFunc("/{fileName}", h.GetFile).Methods("GET")
	fileRouter.HandleFunc("", h.Upload).Methods("POST")
	fileRouter.HandleFunc("", h.DeleteFile).Methods("DELETE")

	accessRouter := router.PathPrefix("/access").Subrouter()
	accessRouter.Use(middleware.AuthSubRouter([]services.Role{services.Manager, services.Editor, services.User, services.Admin}))

	accessRouter.HandleFunc("/{fileName}", h.ServeFile).Methods("GET")

	return router
}
