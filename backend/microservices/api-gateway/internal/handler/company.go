package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// @Summary		Get companies paginated.
// @Tags		Company
// @Produce		json
// @Param		pageNumber   	query     string  false  "Page number"
// @Param 		pageSize    	query     string  false  "Page size"
// @Success		200		{object}	dto.PaginationCompany
// @Failure		400		{string}	string
// @Failure 	401	    {object}	string
// @Router		/api/company [get]
// @Security 	Bearer
func (data *HandlerData) GetCompanies(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("%s/company", data.UrlManagement)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodGet, r.Body)
	functionHandler(w, r)
}

// @Summary		Get company by id.
// @Tags		Company
// @Produce		json
// @Param 		companyId  path string true "Company ID"
// @Success		200		{object}	dto.Company
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Failure		404	    {object}	string
// @Router		/api/company/{companyId} [get]
// @Security 	Bearer
func (data *HandlerData) GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	requestUrl := fmt.Sprintf("%s/company/%s", data.UrlManagement, companyId)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodGet, r.Body)
	functionHandler(w, r)
}

// @Summary		Create company.
// @Tags		Company
// @Accept		json
// @Produce		json
// @Param 		login	body	dto.CompanyData	true "Body"
// @Success		201		{object}	dto.Company
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Router		/api/company [post]
// @Security 	Bearer
func (data *HandlerData) CreateCompany(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("%s/company", data.UrlManagement)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodPost, r.Body)
	functionHandler(w, r)
}

// @Summary		Update company data.
// @Tags		Company
// @Accept		json
// @Produce		json
// @Param 		companyId  path string true "Company ID"
// @Param 		login	body	dto.CompanyData	true "Body"
// @Success		204
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Router		/api/company/{companyId} [patch]
// @Security 	Bearer
func (data *HandlerData) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	requestUrl := fmt.Sprintf("%s/company/%s", data.UrlManagement, companyId)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodPatch, r.Body)
	functionHandler(w, r)
}

// @Summary		Delete company.
// @Tags		Company
// @Produce		json
// @Param 		companyId  path string true "Company ID"
// @Success		204
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Failure		404	    {object}	string
// @Router		/api/company/{companyId} [delete]
// @Security 	Bearer
func (data *HandlerData) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	requestUrl := fmt.Sprintf("%s/company/%s", data.UrlManagement, companyId)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodDelete, r.Body)
	functionHandler(w, r)
}
