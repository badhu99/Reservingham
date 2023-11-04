package server

import (
	// "log"
	// "net/http"

	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/badhu99/api-gateway/docs"
	"github.com/badhu99/api-gateway/internal/handler"
	"github.com/badhu99/api-gateway/internal/middleware"
	"github.com/badhu99/api-gateway/internal/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
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

func (s *Server) AuthRoutes() *mux.Router {

	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatalln("Check .env file: ", err)
	}

	h := handler.HandlerData{
		UrlAuth:       os.Getenv("URL_AUTH"),
		UrlManagement: os.Getenv("URL_MANAGEMENT"),
	}

	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()

	port, err := strconv.Atoi(os.Getenv("PORT"))
	url := fmt.Sprintf("0.0.0.0:%d/swagger/doc.json", port)

	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL(url),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)

	router.Use(middleware.Log)

	routes.AuthRoutes(router, h)
	routes.UserRoutes(router, h)
	routes.CompanyRoutes(router, h)
	routes.PermissionRoutes(router, h)

	return router
}

// func (s *Server) AuthRoutes() *mux.Router {
// 	r := mux.NewRouter()

// 	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
// 		httpSwagger.URL("http://localhost:8081/swagger/doc.json"), //The url pointing to API definition
// 		httpSwagger.DeepLinking(true),
// 		httpSwagger.DocExpansion("none"),
// 		httpSwagger.DomID("swagger-ui"),
// 	)).Methods(http.MethodGet)

// 	return r
// }
