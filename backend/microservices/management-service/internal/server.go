package server

import (
	"log"
	"os"

	"github.com/badhu99/management-service/internal/handler"
	"github.com/badhu99/management-service/internal/middleware"
	"github.com/badhu99/management-service/internal/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
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
		Database: server.Database,
	}

	router.Use(middleware.Log)

	routerCompany := router.PathPrefix("/company").Subrouter()
	routerCompany.Use(middleware.AuthSubRouter([]services.Role{services.Admin}))

	routerUser := router.PathPrefix("/user").Subrouter()
	routerUser.Use(middleware.AuthSubRouter([]services.Role{services.Manager}))

	routerCompany.HandleFunc("", h.GetCompanys).Methods("GET")
	routerCompany.HandleFunc("", h.CreateCompany).Methods("POST")
	routerCompany.HandleFunc("/{companyId}", h.UpdateCompany).Methods("PATCH")
	routerCompany.HandleFunc("/{companyId}", h.DeleteCompany).Methods("DELETE")

	routerUser.HandleFunc("", h.GetUsers).Methods("GET")
	routerUser.HandleFunc("/{userId}", h.GetUserById).Methods("GET")
	routerUser.HandleFunc("", h.CreateUser).Methods("POST")
	routerUser.HandleFunc("/{userId}", h.DeleteUser).Methods("DELETE")
	routerUser.HandleFunc("/{userId}", h.UpdateUser).Methods("PATCH")

	routerUser.HandleFunc("/{userId}/{roleId}", h.AddPermission).Methods("POST")
	routerUser.HandleFunc("/{userId}/{roleId}", h.RemovePermission).Methods("DELETE")
	// TODO
	return router
}

func DatabaseConnect() *gorm.DB {
	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatalln("Check .env file: ", err)
	}

	s := os.Getenv("SQL_DATABASE_URL")

	db, err := gorm.Open(sqlserver.Open(s), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
