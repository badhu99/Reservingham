package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// @Summary		Get users paginated.
// @Tags		User
// @Produce		json
// @Param		pageNumber    query     string  false  "Page number"
// @Param 		pageSize    query     string  false  "Page size"
// @Success		200		{object}	dto.PaginationUsers
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Failure		404	    {object}	string
// @Router		/api/user [get]
// @Security 	Bearer
func (data *HandlerData) GetUsers(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("%s/user", data.UrlManagement)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodGet, r.Body)
	functionHandler(w, r)
}

// @Summary		Get user by id.
// @Tags		User
// @Produce		json
// @Param 		userId  path string true "User ID"
// @Success		200		{object}	dto.UserDataResponse
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Failure		404	    {object}	string
// @Router		/api/user/{userId} [get]
// @Security 	Bearer
func (data *HandlerData) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	requestUrl := fmt.Sprintf("%s/user/%s", data.UrlManagement, userId)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodGet, r.Body)
	functionHandler(w, r)
}

// @Summary		Create user.
// @Tags		User
// @Accept		json
// @Produce		json
// @Param 		login	body	dto.UserData	true "Body"
// @Success		201		{object}	dto.User
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Router		/api/user [post]
// @Security 	Bearer
func (data *HandlerData) CreateUser(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("%s/user", data.UrlManagement)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodPost, r.Body)
	functionHandler(w, r)
}

// @Summary		Update user data.
// @Tags		User
// @Accept		json
// @Produce		json
// @Param 		userId  path string true "User ID"
// @Param 		login	body	dto.UserData	true "Body"
// @Success		204
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Router		/api/user/{userId} [patch]
// @Security 	Bearer
func (data *HandlerData) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	requestUrl := fmt.Sprintf("%s/user/%s", data.UrlManagement, userId)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodPatch, r.Body)
	functionHandler(w, r)
}

// @Summary		Delete user.
// @Tags		User
// @Produce		json
// @Param 		userId  path string true "User ID"
// @Success		204
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Failure		404	    {object}	string
// @Router		/api/user/{userId} [delete]
// @Security 	Bearer
func (data *HandlerData) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]

	requestUrl := fmt.Sprintf("%s/user/%s", data.UrlManagement, userId)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodDelete, r.Body)
	functionHandler(w, r)
}
