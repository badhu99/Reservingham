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
		UrlAuth:                   os.Getenv("URL_AUTH"),
		UrlOrganizationManagement: os.Getenv("URL_MANAGEMENT"),
		UrlContentManagement:      os.Getenv("URL_CONTENT_MANAGEMENT"),
	}

	router := mux.NewRouter()
	router = router.PathPrefix("/api").Subrouter()

	port, _ := strconv.Atoi(os.Getenv("PORT"))
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
	routes.RoleRoutes(router, h)

	routes.DraftRoutes(router, h)

	return router
}
