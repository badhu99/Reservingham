package routes

import (
	"github.com/badhu99/api-gateway/internal/handler"
	"github.com/gorilla/mux"
)

func AuthRoutes(router *mux.Router, handlerData handler.HandlerData) {
	routerCompany := router.PathPrefix("/auth").Subrouter()
	routerCompany.HandleFunc("/signin", handlerData.SignIn).Methods("POST")
}
