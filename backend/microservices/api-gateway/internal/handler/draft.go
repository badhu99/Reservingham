package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func (data *HandlerData) GetDrafts(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("%s/draft", data.UrlContentManagement)
	log.Println(requestUrl)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodGet, r.Body)
	functionHandler(w, r)
}

func (data *HandlerData) CreateDraft(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("%s/draft", data.UrlContentManagement)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodPost, r.Body)
	functionHandler(w, r)
}

func (data *HandlerData) AddDraftHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	draftId := vars["draftId"]

	requestUrl := fmt.Sprintf("%s/draft/%s", data.UrlContentManagement, draftId)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodPost, r.Body)
	functionHandler(w, r)
}
