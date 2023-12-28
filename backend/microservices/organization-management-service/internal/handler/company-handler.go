package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/badhu99/organization-management-service/internal/dto"
	"github.com/badhu99/organization-management-service/internal/utility"

	// mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

func (data *HandlerData) CreateCompany(w http.ResponseWriter, r *http.Request) {

	var companyData dto.CompanyData

	// Validate request
	e, statusCode := utility.ValidateBody(&companyData, r.Body)
	if e != nil {
		http.Error(w, e.Error(), statusCode)
		return
	}

	code, err := data.Services.CreateCompany(companyData)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (data *HandlerData) DeleteCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	companyIdString := vars["companyId"]

	code, err := data.Services.DeleteCompany(companyIdString)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (data *HandlerData) GetCompanies(w http.ResponseWriter, r *http.Request) {
	pageNumber, err := strconv.Atoi(r.URL.Query().Get("pageNumber"))
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		pageSize = 12
	}

	response := data.Services.GetCompanies(pageNumber, pageSize)

	responseJson, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (data *HandlerData) GetCompanyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyIdString := vars["companyId"]

	response, code, err := data.Services.GetCompany(companyIdString)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	responseJson, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (data *HandlerData) UpdateCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	companyIdString := vars["companyId"]

	dataCompany := dto.CompanyData{}
	err, statusCode := utility.ValidateBody(&dataCompany, r.Body)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	code, err := data.Services.UpdateCompany(companyIdString, dataCompany)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
