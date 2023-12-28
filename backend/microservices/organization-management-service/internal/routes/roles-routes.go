package routes

import (
	"github.com/badhu99/organization-management-service/internal/handler"
	"github.com/badhu99/organization-management-service/internal/middleware"
	"github.com/badhu99/organization-management-service/internal/services"
	"github.com/gorilla/mux"
)

func Role(router *mux.Router, handler handler.HandlerData) {
	routerCompany := router.PathPrefix("/role").Subrouter()
	routerCompany.Use(middleware.AuthSubRouter([]services.Role{services.Manager}))

	routerCompany.HandleFunc("", handler.GetRoles).Methods("GET")
}
