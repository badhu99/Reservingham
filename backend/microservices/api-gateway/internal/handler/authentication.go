package handler

import (
	"fmt"
	"net/http"
)

func (data *HandlerData) SignIn(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("%s/signin", data.UrlAuth)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodPost, r.Body)
	functionHandler(w, r)
}
