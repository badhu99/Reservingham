package services

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/badhu99/organization-management-service/internal/dto"
	"github.com/badhu99/organization-management-service/internal/entity"
	"github.com/badhu99/organization-management-service/internal/services"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
)

type BaseServiceData struct {
	Database *gorm.DB
}

func NewBaseServiceData(dbConnection *gorm.DB) *BaseServiceData {
	return &BaseServiceData{
		Database: dbConnection,
	}
}

func (data *BaseServiceData) GetAllUsers(companyIdString string, pageNumber, pageSize int, search string) (dto.Pagination[dto.UserResponse], int, error) {
	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	var count int64
	entityUsers := []entity.User{}

	query := data.Database.Table("[User] as u").Model([]entity.User{}).
		Distinct("u.Id").
		Select("u.*")

	if search != "" {
		query = query.Where("u.Username Like ?", fmt.Sprintf("%%%s%%", search)).
			Or("u.Email Like ?", fmt.Sprintf("%%%s%%", search))
	}

	query.
		Joins("JOIN Permission p ON u.Id = p.UserId").
		Preload("Permissions").
		Not("p.CompanyId IS NULL").
		Where("p.CompanyId = ?", companyId).
		Joins("JOIN Role r ON r.id = p.RoleId").
		Preload("Permissions.Role").
		Offset((pageNumber - 1) * pageSize).Limit(pageSize).
		Find(&entityUsers)

	data.Database.Table("Permission as p").
		Distinct("p.UserId").
		Where("p.CompanyId = ?", companyId).
		Count(&count)

	responseUser := []dto.UserResponse{}
	for _, perm := range entityUsers {
		roles := []dto.RoleResponse{}
		for _, r := range perm.Permissions {
			if r.Role.Level < 4 {
				roles = append(roles, dto.RoleResponse{
					ID:   r.RoleID,
					Name: r.Role.Name,
				})
			}
		}
		responseUser = append(responseUser, dto.UserResponse{
			User: dto.User{
				ID: perm.ID,
				UserData: dto.UserData{
					Email:     perm.Email,
					Username:  perm.Username,
					Firstname: perm.Firstname,
					Lastname:  perm.Lastname,
				},
			},
			Roles: roles,
		})
	}

	return dto.Pagination[dto.UserResponse]{
		Count: int(count),
		Page:  pageNumber,
		Size:  pageSize,
		Items: responseUser,
	}, 0, nil
}

func (data *BaseServiceData) GetUserById(userIdString, companyIdString string) (dto.UserResponse, int, error) {

	userId := mssql.UniqueIdentifier{}
	userId.UnmarshalJSON([]byte(userIdString))

	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	entityUser := entity.User{ID: userId}
	errorGorm := data.Database.
		Joins("JOIN Permission p ON [User].Id = p.UserId").
		Preload("Permissions").
		Where("p.CompanyId = ?", companyId).
		Joins("JOIN Role r ON r.id = p.RoleId").
		Preload("Permissions.Role").
		First(&entityUser).Error

	if errorGorm == gorm.ErrRecordNotFound {
		return dto.UserResponse{}, http.StatusNotFound, fmt.Errorf(http.StatusText(http.StatusNotFound))
	}
	if errorGorm != nil {
		return dto.UserResponse{}, http.StatusInternalServerError, errorGorm
	}

	responseRoles := []dto.RoleResponse{}
	for _, r := range entityUser.Permissions {
		if r.Role.Level < 4 {
			responseRoles = append(responseRoles, dto.RoleResponse{
				ID:   r.RoleID,
				Name: r.Role.Name,
			})
		}
	}

	return dto.UserResponse{
		User: dto.User{
			ID: entityUser.ID,
			UserData: dto.UserData{
				Email:     entityUser.Email,
				Username:  entityUser.Username,
				Firstname: entityUser.Firstname,
				Lastname:  entityUser.Lastname,
			},
		},
		Roles: responseRoles,
	}, 0, nil
}

func (data *BaseServiceData) InviteUser(companyIdString, email string) error {

	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	entityRole := entity.Role{Name: "User"}
	errorGorm := data.Database.First(&entityRole, "Name = ?", entityRole.Name).Error
	if errorGorm != nil {
		return errorGorm
	}

	entityUser := entity.User{}
	errorGorm = data.Database.Where("Email = ?", email).First(&entityUser).Error
	if errorGorm == gorm.ErrRecordNotFound {
		entityUser.Email = email
		randomString, _ := services.GenerateRandomString(20)
		entityUser.Code = randomString
		entityUser.PasswordChange = true
		data.Database.Create(&entityUser)
	}

	entityPermission := entity.Permission{}
	errorGorm = data.Database.Where("UserId = ? AND RoleId = ? AND CompanyId = ?", entityUser.ID, entityRole.ID, companyId).First(&entityPermission).Error
	if errorGorm == gorm.ErrRecordNotFound {
		entityPermission = entity.Permission{
			UserID:    entityUser.ID,
			RoleID:    entityRole.ID,
			CompanyID: companyId,
		}
		data.Database.Create(&entityPermission)
	}

	return nil
}

func (data *BaseServiceData) CreateUser(userData dto.UserInviteConfirm) (int, error) {

	entityUser := entity.User{}
	errorGorm := data.Database.
		Where("Email = ?", userData.Email).
		First(&entityUser).
		Error

	if errorGorm == gorm.ErrRecordNotFound {
		return http.StatusBadRequest, fmt.Errorf("user does not exists")
	}

	if entityUser.Code != userData.Code || !entityUser.PasswordChange {
		return http.StatusBadRequest, fmt.Errorf("code does not match")
	}

	// Basic user data
	entityUser.Firstname = userData.Firstname
	entityUser.Lastname = userData.Lastname
	entityUser.Username = userData.Username

	// Password user data
	salt, _ := services.GenerateRandomSalt(16)
	password, _ := services.HashPassword([]byte(userData.Password), salt)
	entityUser.PasswordHash = base64.StdEncoding.EncodeToString(password)
	entityUser.Salt = base64.StdEncoding.EncodeToString(salt)

	// Disable re-login
	entityUser.PasswordChange = false
	entityUser.Code = ""

	errorGorm = data.Database.
		Where("Username = ?", entityUser.Username).
		First(&entity.User{}).
		Error

	if errorGorm != gorm.ErrRecordNotFound {
		return http.StatusBadRequest, fmt.Errorf("username already taken")
	}

	errGorm := data.Database.Save(&entityUser).Error
	if errGorm != nil {
		return http.StatusInternalServerError, errGorm
	}
	return 0, nil
}

func (data *BaseServiceData) DeleteUser(companyIdString, userIdString string) (int, error) {
	userId := mssql.UniqueIdentifier{}
	companyId := mssql.UniqueIdentifier{}

	userId.UnmarshalJSON([]byte(userIdString))
	companyId.UnmarshalJSON([]byte(companyIdString))

	entityUser := entity.User{ID: userId}
	errorGorm := data.Database.
		Joins("JOIN Permission p ON p.UserId = [User].Id").
		Preload("Permissions").
		Where("p.CompanyId = ?", companyId).
		First(&entityUser).
		Error

	if errorGorm == gorm.ErrRecordNotFound {
		return http.StatusNotFound, fmt.Errorf(http.StatusText(http.StatusNotFound))
	} else if errorGorm != nil {
		return http.StatusInternalServerError, errorGorm
	}

	data.Database.Delete(&entityUser)
	return 0, nil
}

func (data *BaseServiceData) UpdateUser(companyIdString, userIdString string, userData dto.UserData) (int, error) {

	userId := mssql.UniqueIdentifier{}
	companyId := mssql.UniqueIdentifier{}

	userId.UnmarshalJSON([]byte(userIdString))
	companyId.UnmarshalJSON([]byte(companyIdString))

	entityUser := entity.User{ID: userId}
	errorGorm := data.Database.
		Joins("JOIN Permission p ON p.UserId = [User].Id").
		Preload("Permissions").
		Where("p.CompanyId = ?", companyId).
		First(&entityUser).
		Error

	if errorGorm == gorm.ErrRecordNotFound {
		return http.StatusNotFound, fmt.Errorf(http.StatusText(http.StatusNotFound))
	} else if errorGorm != nil {
		return http.StatusInternalServerError, errorGorm
	}

	if userData.Email != "" && userData.Email != entityUser.Email {
		entityUser.Email = userData.Email
	}

	if userData.Username != "" && userData.Username != entityUser.Username {
		entityUser.Username = userData.Username
	}

	if userData.Password != "" {
		saltBytes, _ := services.GenerateRandomSalt(16)
		password, _ := services.HashPassword([]byte(userData.Password), saltBytes)

		entityUser.PasswordHash = base64.StdEncoding.EncodeToString(password)
		entityUser.Salt = base64.StdEncoding.EncodeToString(saltBytes)
	}

	data.Database.Save(&entityUser)
	return 0, nil
}
