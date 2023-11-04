package routes

import (
	"github.com/badhu99/management-service/internal/handler"
	"github.com/badhu99/management-service/internal/middleware"
	"github.com/badhu99/management-service/internal/services"
	"github.com/gorilla/mux"
)

func Company(router *mux.Router, handler handler.HandlerData) {
	routerCompany := router.PathPrefix("/company").Subrouter()
	routerCompany.Use(middleware.AuthSubRouter([]services.Role{services.Admin}))

	routerCompany.HandleFunc("", handler.GetCompanies).Methods("GET")
	routerCompany.HandleFunc("/{companyId}", handler.GetCompanyById).Methods("GET")
	routerCompany.HandleFunc("", handler.CreateCompany).Methods("POST")
	routerCompany.HandleFunc("/{companyId}", handler.UpdateCompany).Methods("PATCH")
	routerCompany.HandleFunc("/{companyId}", handler.DeleteCompany).Methods("DELETE")
}
