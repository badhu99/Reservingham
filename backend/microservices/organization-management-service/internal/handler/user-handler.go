package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/badhu99/organization-management-service/internal/dto"
	"github.com/badhu99/organization-management-service/internal/utility"
	"github.com/gorilla/mux"
)

func (data *HandlerData) GetUsers(w http.ResponseWriter, r *http.Request) {

	companyIdString, ok := r.Context().Value("companyId").(string)
	if !ok {
		http.Error(w, "CompanyId not found", http.StatusBadRequest)
		return
	}

	pageNumber, err := strconv.Atoi(r.URL.Query().Get("pageNumber"))
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		pageSize = 12
	}

	searchParam := r.URL.Query().Get("search")

	response, code, err := data.Services.GetAllUsers(companyIdString, pageNumber, pageSize, searchParam)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}
	responseJson, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(responseJson))
}

func (data *HandlerData) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdString := vars["userId"]

	companyIdString, _ := r.Context().Value("companyId").(string)

	response, code, err := data.Services.GetUserById(userIdString, companyIdString)
	if err == nil {
		http.Error(w, err.Error(), code)
		return
	}

	responseJson, e := json.Marshal(response)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (data *HandlerData) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdString := vars["userId"]

	companyIdString, _ := r.Context().Value("companyId").(string)

	code, err := data.Services.DeleteUser(userIdString, companyIdString)
	if err == nil {
		http.Error(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (data *HandlerData) UpdateUser(w http.ResponseWriter, r *http.Request) {

	dataUser := dto.UserData{}
	err, code := utility.ValidateBody(&dataUser, r.Body)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	vars := mux.Vars(r)
	userIdString := vars["userId"]

	companyIdString, _ := r.Context().Value("companyId").(string)

	code, err = data.Services.UpdateUser(companyIdString, userIdString, dataUser)
	if err == nil {
		http.Error(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (data *HandlerData) AddPermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdString := vars["userId"]
	roleIdString := vars["roleId"]

	companyIdString, _ := r.Context().Value("companyId").(string)

	code, err := data.Services.AddPermissions(userIdString, companyIdString, roleIdString)
	if err == nil {
		http.Error(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (data *HandlerData) RemovePermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdString := vars["userId"]
	roleIdString := vars["roleId"]

	companyIdString, _ := r.Context().Value("companyId").(string)

	code, err := data.Services.RemovePermissions(userIdString, companyIdString, roleIdString)
	if err == nil {
		http.Error(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (data *HandlerData) InviteUser(w http.ResponseWriter, r *http.Request) {
	companyIdString, ok := r.Context().Value("companyId").(string)
	if !ok {
		http.Error(w, "UserId not found", http.StatusBadRequest)
		return
	}

	userInviteDto := dto.UserInvite{}
	err, code := utility.ValidateBody(&userInviteDto, r.Body)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	err = data.Services.InviteUser(companyIdString, userInviteDto.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (data *HandlerData) CreateUser(w http.ResponseWriter, r *http.Request) {

	var dataUser dto.UserInviteConfirm
	e, statusCode := utility.ValidateBody(&dataUser, r.Body)
	if e != nil {
		http.Error(w, http.StatusText(statusCode), statusCode)
		return
	}

	code, err := data.Services.CreateUser(dataUser)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
