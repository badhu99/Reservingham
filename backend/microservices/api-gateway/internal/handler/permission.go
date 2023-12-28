package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// @Summary		Add user permission.
// @Tags		Permission
// @Param 		userId  path string true "UserId ID"
// @Param 		roleId  path string true "Role ID"
// @Success		201
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Failure		404	    {object}	string
// @Router		/api/permission/{userId}/{roleId} [post]
// @Security 	Bearer
func (data *HandlerData) AddPermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	roleId := vars["roleId"]

	requestUrl := fmt.Sprintf("%s/user/%s/%s", data.UrlOrganizationManagement, userId, roleId)
	log.Println(requestUrl)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodPost, r.Body)
	functionHandler(w, r)
}

// @Summary		Delete user permission.
// @Tags		Permission
// @Param 		userId  path string true "UserId ID"
// @Param 		roleId  path string true "Role ID"
// @Success		204
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Failure		404	    {object}	string
// @Router		/api/permission/{userId}/{roleId} [delete]
// @Security 	Bearer
func (data *HandlerData) DeletePermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	roleId := vars["roleId"]

	requestUrl := fmt.Sprintf("%s/user/%s/%s", data.UrlOrganizationManagement, userId, roleId)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodDelete, r.Body)
	functionHandler(w, r)
}
