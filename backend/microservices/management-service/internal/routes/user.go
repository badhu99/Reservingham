package routes

import (
	"github.com/badhu99/management-service/internal/handler"
	"github.com/badhu99/management-service/internal/middleware"
	"github.com/badhu99/management-service/internal/services"
	"github.com/gorilla/mux"
)

func User(router *mux.Router, handler handler.HandlerData) {
	routerUser := router.PathPrefix("/user").Subrouter()
	routerUser.Use(middleware.AuthSubRouter([]services.Role{services.Manager}))

	routerUser.HandleFunc("", handler.GetUsers).Methods("GET")
	routerUser.HandleFunc("/{userId}", handler.GetUserById).Methods("GET")
	routerUser.HandleFunc("", handler.CreateUser).Methods("POST")
	routerUser.HandleFunc("/{userId}", handler.DeleteUser).Methods("DELETE")
	routerUser.HandleFunc("/{userId}", handler.UpdateUser).Methods("PATCH")

	routerUser.HandleFunc("/{userId}/{roleId}", handler.AddPermission).Methods("POST")
	routerUser.HandleFunc("/{userId}/{roleId}", handler.RemovePermission).Methods("DELETE")
}
