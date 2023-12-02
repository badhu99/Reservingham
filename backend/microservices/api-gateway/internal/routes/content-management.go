package routes

import (
	"github.com/badhu99/api-gateway/internal/handler"
	"github.com/gorilla/mux"
)

func DraftRoutes(router *mux.Router, handlerData handler.HandlerData) {
	routerUser := router.PathPrefix("/draft").Subrouter()
	routerUser.HandleFunc("", handlerData.GetDrafts).Methods("GET")
	routerUser.HandleFunc("", handlerData.CreateDraft).Methods("POST")
	routerUser.HandleFunc("/{draftId}", handlerData.AddDraftHistory).Methods("POST")
}
