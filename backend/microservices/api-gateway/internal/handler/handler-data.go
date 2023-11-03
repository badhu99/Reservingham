package handler

import (
	"io"
	"net/http"
)

type HandlerData struct {
	UrlAuth       string
	UrlManagement string
}

func (data *HandlerData) HttpRequestBroker(requestUrl, method string, body io.ReadCloser) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		request, _ := http.NewRequest(method, requestUrl, body)
		request.Header.Add("Content-Type", "application/json")

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		defer response.Body.Close()
		responseBody, _ := io.ReadAll(response.Body)

		w.Header().Set("Content-Type", response.Header.Get("Content-Type"))
		w.WriteHeader(response.StatusCode)
		w.Write(responseBody)
	}
}
