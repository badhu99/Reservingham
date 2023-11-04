package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HandlerData struct {
	UrlAuth       string
	UrlManagement string
}

func (data *HandlerData) HttpRequestBroker(requestUrl, method string, body io.ReadCloser) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		queryParams := url.Values{}
		for paramName, paramValues := range r.URL.Query() {
			for _, singleParamvalue := range paramValues {
				queryParams.Add(paramName, singleParamvalue)
			}
		}
		requestUrl = fmt.Sprintf("%s?%s", requestUrl, queryParams.Encode())

		request, _ := http.NewRequest(method, requestUrl, body)
		request.Header.Add("Content-Type", "application/json")
		auth := r.Header.Get("Authorization")
		if auth != "" {
			request.Header.Add("Authorization", auth)
		}

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

func (data *HandlerData) Test(w http.ResponseWriter, r *http.Request) {
	spec, err := ioutil.ReadFile("docs/swagger.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(spec)
}
