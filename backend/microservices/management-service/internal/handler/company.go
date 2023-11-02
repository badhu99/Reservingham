package handler

import (
	"fmt"
	"net/http"

	"github.com/badhu99/management-service/internal/dto"
	"github.com/badhu99/management-service/internal/entity"
	"github.com/badhu99/management-service/internal/utility"
	mssql "github.com/microsoft/go-mssqldb"

	// mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (base *HandlerData) CreateCompany(w http.ResponseWriter, r *http.Request) {

	var companyData dto.CompanyData

	// Validate request
	e, statusCode := utility.ValidateBody(&companyData, r.Body)
	if e != nil {
		http.Error(w, e.Error(), statusCode)
		return
	}

	entityCompany := entity.Company{
		Name: companyData.Name,
	}

	// Check if company already exists
	errGorm := base.Database.First(&entityCompany, "Name = ?", entityCompany.Name).Error

	if errGorm != gorm.ErrRecordNotFound {
		http.Error(w, fmt.Sprintf("Company with name %s already exists", companyData.Name), http.StatusBadRequest)
		return
	}

	base.Database.Create(&entityCompany)

	w.WriteHeader(http.StatusAccepted)
}

func (base *HandlerData) DeleteCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	companyIdString := vars["companyId"]

	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	entityCompany := entity.Company{
		ID: companyId,
	}

	errorGorm := base.Database.First(&entityCompany).Error
	if errorGorm == gorm.ErrRecordNotFound {
		http.Error(w, fmt.Sprintf("Company with id %s not found!", companyId.String()), http.StatusBadRequest)
		return
	} else if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
	}

	base.Database.Delete(&entityCompany)

	w.WriteHeader(http.StatusAccepted)
}

func (data *HandlerData) GetCompanys(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get company"))
}

func (data *HandlerData) GetCompanyById(w http.ResponseWriter, r *http.Request) {

}

func (base *HandlerData) UpdateCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	companyId := vars["companyId"]

	w.Write([]byte(companyId))
}
