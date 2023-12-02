package server

import (
	"log"
	"os"

	"github.com/badhu99/content-management-service/internal/handler"
	"github.com/badhu99/content-management-service/internal/middleware"
	"github.com/badhu99/content-management-service/internal/routes"
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
		Database:                 DatabaseConnect(),
		UrlServiceUserManagement: "http://localhost:8082",
	}

	router.Use(middleware.Log)

	routes.Draft(router, h)

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
