package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/badhu99/api-gateway/internal/dto"
	"github.com/badhu99/api-gateway/internal/utility"
)

func (data *HandlerData) SignIn(w http.ResponseWriter, r *http.Request) {

	dataLogin := dto.Login{}

	e, statusCode := utility.ValidateBody(&dataLogin, r.Body)
	if e != nil {
		http.Error(w, e.Error(), statusCode)
		return
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(dataLogin)
	requestUrl := fmt.Sprintf("%s/signin", data.UrlAuth)

	request, _ := http.NewRequest(http.MethodPost, requestUrl, body)
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
	w.Write(responseBody)
	w.WriteHeader(response.StatusCode)
}
