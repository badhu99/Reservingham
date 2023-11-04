package routes

import (
	"github.com/badhu99/api-gateway/internal/handler"
	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router, handlerData handler.HandlerData) {
	routerUser := router.PathPrefix("/user").Subrouter()
	routerUser.HandleFunc("", handlerData.GetUsers).Methods("GET")
	routerUser.HandleFunc("/{userId}", handlerData.GetUser).Methods("GET")
	routerUser.HandleFunc("", handlerData.CreateUser).Methods("POST")
	routerUser.HandleFunc("/{userId}", handlerData.UpdateUser).Methods("PATCH")
	routerUser.HandleFunc("/{userId}", handlerData.DeleteUser).Methods("DELETE")
}

func CompanyRoutes(router *mux.Router, handlerData handler.HandlerData) {
	routerCompany := router.PathPrefix("/company").Subrouter()
	routerCompany.HandleFunc("", handlerData.GetCompanies).Methods("GET")
	routerCompany.HandleFunc("/{companyId}", handlerData.GetCompany).Methods("GET")
	routerCompany.HandleFunc("", handlerData.CreateCompany).Methods("POST")
	routerCompany.HandleFunc("/{companyId}", handlerData.UpdateCompany).Methods("PATCH")
	routerCompany.HandleFunc("/{companyId}", handlerData.DeleteCompany).Methods("DELETE")
}

func PermissionRoutes(router *mux.Router, handlerData handler.HandlerData) {
	routerPermissions := router.PathPrefix("/permission").Subrouter()
	routerPermissions.HandleFunc("/{userId}/{roleId}", handlerData.AddPermission).Methods("POST")
	routerPermissions.HandleFunc("/{userId}/{roleId}", handlerData.DeletePermission).Methods("DELETE")
}

func RoleRoutes(router *mux.Router, handlerData handler.HandlerData) {
	routerPermissions := router.PathPrefix("/role").Subrouter()
	routerPermissions.HandleFunc("", handlerData.GetRoles).Methods("GET")
}
