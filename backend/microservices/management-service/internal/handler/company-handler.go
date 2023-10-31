package handler

import "net/http"

func (base *HandlerData) CreateCompany(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It works"))
}
