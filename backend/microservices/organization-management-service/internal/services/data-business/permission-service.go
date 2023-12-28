package services

import (
	"fmt"
	"net/http"

	"github.com/badhu99/organization-management-service/internal/entity"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
)

func (data *BaseServiceData) AddPermissions(userIdString, companyIdString, roleIdString string) (int, error) {
	userId := mssql.UniqueIdentifier{}
	companyId := mssql.UniqueIdentifier{}
	roleId := mssql.UniqueIdentifier{}

	companyId.UnmarshalJSON([]byte(companyIdString))
	userId.UnmarshalJSON([]byte(userIdString))
	roleId.UnmarshalJSON([]byte(roleIdString))

	entityRole := entity.Role{}
	errorGorm := data.Database.
		Where("Id = ? AND Level < 4", roleId).
		First(&entityRole).
		Error

	if errorGorm == gorm.ErrRecordNotFound {
		return http.StatusNotFound, fmt.Errorf(http.StatusText(http.StatusNotFound))
	} else if errorGorm != nil {
		return http.StatusInternalServerError, errorGorm
	}

	entityPermission := entity.Permission{}
	errorGorm = data.Database.Model(entity.Permission{}).
		Where("CompanyId = ? AND UserId = ? AND RoleId = ?", companyId, userId, roleId).
		First(&entityPermission).
		Error

	if errorGorm != gorm.ErrRecordNotFound {
		return http.StatusBadRequest, fmt.Errorf("Permissions already set")
	}

	entityPermission.CompanyID = companyId
	entityPermission.UserID = userId
	entityPermission.RoleID = roleId

	errorGorm = data.Database.Create((&entityPermission)).Error

	if errorGorm == gorm.ErrRecordNotFound {
		return http.StatusNotFound, fmt.Errorf(http.StatusText(http.StatusNotFound))
	} else if errorGorm != nil {
		return http.StatusInternalServerError, errorGorm
	}
	return 0, nil
}

func (data *BaseServiceData) RemovePermissions(userIdString, companyIdString, roleIdString string) (int, error) {
	userId := mssql.UniqueIdentifier{}
	companyId := mssql.UniqueIdentifier{}
	roleId := mssql.UniqueIdentifier{}

	companyId.UnmarshalJSON([]byte(companyIdString))
	userId.UnmarshalJSON([]byte(userIdString))
	roleId.UnmarshalJSON([]byte(roleIdString))

	entityPermission := entity.Permission{}
	errorGorm := data.Database.
		Where("CompanyId = ? AND UserId = ? AND RoleId = ?", companyId, userId, roleId).
		First(&entityPermission).
		Error

	if errorGorm == gorm.ErrRecordNotFound {
		return http.StatusNotFound, fmt.Errorf(http.StatusText(http.StatusNotFound))
	} else if errorGorm != nil {
		return http.StatusInternalServerError, errorGorm
	}

	data.Database.Delete(&entityPermission)
	return 0, nil
}
