package services

import (
	"fmt"
	"net/http"

	"github.com/badhu99/organization-management-service/internal/dto"
	"github.com/badhu99/organization-management-service/internal/entity"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
)

func (data *BaseServiceData) CreateCompany(companyData dto.CompanyData) (int, error) {

	entityCompany := entity.Company{
		Name: companyData.Name,
	}

	// Check if company already exists
	errGorm := data.Database.First(&entityCompany, "Name = ?", entityCompany.Name).Error

	if errGorm != gorm.ErrRecordNotFound {
		return http.StatusBadRequest, fmt.Errorf("Company with name %s already exists", companyData.Name)
	}

	data.Database.Create(&entityCompany)

	return 0, nil
}

func (data *BaseServiceData) UpdateCompany(companyIdString string, dataCompany dto.CompanyData) (int, error) {
	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))
	entityCompany := entity.Company{ID: companyId}

	errorGorm := data.Database.Model(&entityCompany).First(&entityCompany).Error

	if errorGorm == gorm.ErrRecordNotFound {
		return http.StatusNotFound, fmt.Errorf(http.StatusText(http.StatusNotFound))
	} else if errorGorm != nil {
		return http.StatusInternalServerError, errorGorm
	}

	gormError := data.Database.First(&entity.Company{}, "Name = ?", dataCompany.Name).Error
	if gormError != gorm.ErrRecordNotFound {
		return http.StatusBadRequest, fmt.Errorf("Name already exists")
	}

	if dataCompany.Name != "" {
		entityCompany.Name = dataCompany.Name
	}

	data.Database.Save(entityCompany)
	return 0, nil
}

func (data *BaseServiceData) DeleteCompany(companyIdString string) (int, error) {
	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	entityCompany := entity.Company{
		ID: companyId,
	}

	errorGorm := data.Database.First(&entityCompany).Error
	if errorGorm == gorm.ErrRecordNotFound {
		return http.StatusBadRequest, fmt.Errorf("company with id %s not found", companyId.String())
	} else if errorGorm != nil {
		return http.StatusInternalServerError, errorGorm
	}

	data.Database.Delete(&entityCompany)

	return 0, nil
}

func (data *BaseServiceData) GetCompanies(pageNumber, pageSize int) dto.Pagination[dto.Company] {
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

	return dto.Pagination[dto.Company]{
		Count: int(count),
		Page:  pageNumber,
		Size:  pageSize,
		Items: responseCompany,
	}
}

func (data *BaseServiceData) GetCompany(companyIdString string) (dto.Company, int, error) {
	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	entityCompany := entity.Company{ID: companyId}

	errorGorm := data.Database.Model(&entityCompany).First(&entityCompany).Error

	if errorGorm == gorm.ErrRecordNotFound {
		return dto.Company{}, http.StatusNotFound, fmt.Errorf(http.StatusText(http.StatusNotFound))
	} else if errorGorm != nil {
		return dto.Company{}, http.StatusInternalServerError, errorGorm
	}

	return dto.Company{
		ID: entityCompany.ID,
		CompanyData: dto.CompanyData{
			Name: entityCompany.Name,
		},
	}, 0, nil
}
