package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/badhu99/management-service/internal/dto"
	"github.com/badhu99/management-service/internal/entity"
	"github.com/badhu99/management-service/internal/utility"
	mssql "github.com/microsoft/go-mssqldb"

	// mssql "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func (data *HandlerData) CreateCompany(w http.ResponseWriter, r *http.Request) {

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
	errGorm := data.Database.First(&entityCompany, "Name = ?", entityCompany.Name).Error

	if errGorm != gorm.ErrRecordNotFound {
		http.Error(w, fmt.Sprintf("Company with name %s already exists", companyData.Name), http.StatusBadRequest)
		return
	}

	data.Database.Create(&entityCompany)

	w.WriteHeader(http.StatusAccepted)
}

func (data *HandlerData) DeleteCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	companyIdString := vars["companyId"]

	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	entityCompany := entity.Company{
		ID: companyId,
	}

	errorGorm := data.Database.First(&entityCompany).Error
	if errorGorm == gorm.ErrRecordNotFound {
		http.Error(w, fmt.Sprintf("Company with id %s not found!", companyId.String()), http.StatusBadRequest)
		return
	} else if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
	}

	data.Database.Delete(&entityCompany)

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

	var count int64
	entityCompanys := []entity.Company{}

	data.Database.Model([]entity.Company{}).
		Count(&count).
		Offset((pageNumber - 1) * pageSize).Limit(pageSize).
		Find(&entityCompanys)

	responseCompany := []dto.Company{}
	for _, c := range entityCompanys {
		responseCompany = append(responseCompany, dto.Company{
			ID: c.ID,
			CompanyData: dto.CompanyData{
				Name: c.Name,
			},
		})
	}

	responseData := dto.Pagination[dto.Company]{
		Count: int(count),
		Page:  pageNumber,
		Size:  pageSize,
		Items: responseCompany,
	}

	response, _ := json.Marshal(responseData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (data *HandlerData) GetCompanyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyIdString := vars["companyId"]

	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	entityCompany := entity.Company{ID: companyId}

	errorGorm := data.Database.Model(&entityCompany).First(&entityCompany).Error

	if errorGorm == gorm.ErrRecordNotFound {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	} else if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
	}

	responseCompany := dto.Company{
		ID: entityCompany.ID,
		CompanyData: dto.CompanyData{
			Name: entityCompany.Name,
		},
	}

	response, _ := json.Marshal(responseCompany)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (data *HandlerData) UpdateCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	companyIdString := vars["companyId"]

	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	entityCompany := entity.Company{ID: companyId}

	errorGorm := data.Database.Model(&entityCompany).First(&entityCompany).Error

	if errorGorm == gorm.ErrRecordNotFound {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	} else if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
	}

	dataBody := dto.CompanyData{}
	err, statusCode := utility.ValidateBody(&dataBody, r.Body)
	if err != nil {
		http.Error(w, err.Error(), statusCode)
		return
	}

	gormError := data.Database.First(&entity.Company{}, "Name = ?", dataBody.Name).Error
	if gormError != gorm.ErrRecordNotFound {
		http.Error(w, "Name already exists", http.StatusBadRequest)
		return
	}

	if dataBody.Name != "" {
		entityCompany.Name = dataBody.Name
	}

	data.Database.Save(entityCompany)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

}

func (base *HandlerData) AssignUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyId := vars["companyId"]

	w.Write([]byte(companyId))
}
