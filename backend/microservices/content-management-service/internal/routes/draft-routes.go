package routes

import (
	"github.com/badhu99/content-management-service/internal/handler"
	"github.com/badhu99/content-management-service/internal/middleware"
	"github.com/badhu99/content-management-service/internal/services"
	"github.com/gorilla/mux"
)

func Draft(router *mux.Router, handler handler.HandlerData) {
	routerDraft := router.PathPrefix("/draft").Subrouter()
	routerDraft.Use(middleware.AuthSubRouter([]services.Role{services.Editor}))

	routerDraft.HandleFunc("", handler.GetDrafts).Methods("GET")
	routerDraft.HandleFunc("", handler.CreateDraft).Methods("POST")
	routerDraft.HandleFunc("/{draftId}", handler.AddDraftHistory).Methods("POST")
	routerDraft.HandleFunc("/{draftId}", handler.UpdateDraftData).Methods("PATCH")
}
