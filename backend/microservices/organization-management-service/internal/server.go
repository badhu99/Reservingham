package server

import (
	"github.com/badhu99/organization-management-service/internal/handler"
	"github.com/badhu99/organization-management-service/internal/middleware"
	"github.com/badhu99/organization-management-service/internal/routes"
	services "github.com/badhu99/organization-management-service/internal/services/data-business"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Database *gorm.DB
	Config   Config
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

func (server *Server) Routes() *mux.Router {
	router := mux.NewRouter()

	h := handler.HandlerData{
		Services: services.NewBaseServiceData(server.Database),
	}

	router.Use(middleware.Log)

	routes.Company(router, h)
	routes.User(router, h)
	routes.Role(router, h)

	return router
}
