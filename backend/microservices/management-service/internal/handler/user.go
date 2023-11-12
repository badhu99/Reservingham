package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/badhu99/management-service/internal/dto"
	"github.com/badhu99/management-service/internal/entity"
	"github.com/badhu99/management-service/internal/services"
	"github.com/badhu99/management-service/internal/utility"
	"github.com/gorilla/mux"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
)

func (data *HandlerData) GetUsers(w http.ResponseWriter, r *http.Request) {

	companyIdString, ok := r.Context().Value("companyId").(string)
	if !ok {
		http.Error(w, "UserId not found", http.StatusBadRequest)
		return
	}
	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	pageNumber, err := strconv.Atoi(r.URL.Query().Get("pageNumber"))
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		pageSize = 12
	}

	searchParam := r.URL.Query().Get("search")
	var count int64

	entityUsers := []entity.User{}

	query := data.Database.Table("[User] as u").Model([]entity.User{}).
		Distinct("u.Id").
		Select("u.*")

	if searchParam != "" {
		query = query.Where("u.Username Like ?", fmt.Sprintf("%%%s%%", searchParam)).
			Or("u.Email Like ?", fmt.Sprintf("%%%s%%", searchParam))
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

	responsePaginated := dto.Pagination[dto.UserResponse]{
		Count: int(count),
		Page:  pageNumber,
		Size:  pageSize,
		Items: responseUser,
	}
	response, _ := json.Marshal(responsePaginated)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func (data *HandlerData) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdString := vars["userId"]
	userId := mssql.UniqueIdentifier{}
	userId.UnmarshalJSON([]byte(userIdString))

	companyIdString, _ := r.Context().Value("companyId").(string)
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
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
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

	responseUser := dto.UserResponse{
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
	}

	responseJson, e := json.Marshal(responseUser)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}

func (data *HandlerData) CreateUser(w http.ResponseWriter, r *http.Request) {

	companyIdString, ok := r.Context().Value("companyId").(string)
	if !ok {
		http.Error(w, "UserId not found", http.StatusBadRequest)
		return
	}
	companyId := mssql.UniqueIdentifier{}
	companyId.UnmarshalJSON([]byte(companyIdString))

	var dataUser dto.UserData
	e, statusCode := utility.ValidateBody(&dataUser, r.Body)
	if e != nil {
		http.Error(w, http.StatusText(statusCode), statusCode)
	}

	entityRole := entity.Role{Name: "User"}
	errorGorm := data.Database.First(&entityRole, "Name = ?", entityRole.Name).Error
	if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
	}

	entityUser := entity.User{Username: dataUser.Username}
	errorGorm = data.Database.First(&entityUser, "Username = ?", entityUser.Username).Error
	if errorGorm == gorm.ErrRecordNotFound {
		entityUser.Username = dataUser.Username
		entityUser.Email = dataUser.Email
		entityUser.Firstname = dataUser.Firstname
		entityUser.Lastname = dataUser.Lastname

		salt, _ := services.GenerateRandomSalt(16)
		password, _ := services.HashPassword([]byte(dataUser.Password), salt)
		entityUser.PasswordHash = base64.StdEncoding.EncodeToString(password)
		entityUser.Salt = base64.StdEncoding.EncodeToString(salt)

		data.Database.Create(&entityUser)
	}

	entityPermission := entity.Permission{
		UserID:    entityUser.ID,
		RoleID:    entityRole.ID,
		CompanyID: companyId,
	}
	data.Database.Create(&entityPermission)

	responseUser := dto.User{
		ID: entityUser.ID,
		UserData: dto.UserData{
			Email:     entityUser.Email,
			Username:  entityUser.Username,
			Firstname: entityUser.Firstname,
			Lastname:  entityUser.Lastname,
		},
	}

	response, _ := json.Marshal(responseUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func (data *HandlerData) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdString := vars["userId"]

	companyIdString, _ := r.Context().Value("companyId").(string)

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
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
	}

	data.Database.Delete(&entityUser)

	w.WriteHeader(http.StatusNoContent)
}

func (data *HandlerData) UpdateUser(w http.ResponseWriter, r *http.Request) {

	dataUser := dto.UserData{}
	utility.ValidateBody(&dataUser, r.Body)

	vars := mux.Vars(r)
	userIdString := vars["userId"]

	companyIdString, _ := r.Context().Value("companyId").(string)

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
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
	}

	if dataUser.Email != "" && dataUser.Email != entityUser.Email {
		entityUser.Email = dataUser.Email
	}

	if dataUser.Username != "" && dataUser.Username != entityUser.Username {
		entityUser.Username = dataUser.Username
	}

	if dataUser.Password != "" {
		saltBytes, _ := services.GenerateRandomSalt(16)
		password, _ := services.HashPassword([]byte(dataUser.Password), saltBytes)

		entityUser.PasswordHash = base64.StdEncoding.EncodeToString(password)
		entityUser.Salt = base64.StdEncoding.EncodeToString(saltBytes)
	}

	data.Database.Save(&entityUser)

	w.WriteHeader(http.StatusNoContent)
}

func (data *HandlerData) AddPermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdString := vars["userId"]
	roleIdString := vars["roleId"]

	companyIdString, _ := r.Context().Value("companyId").(string)

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
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
	}

	entityPermission := entity.Permission{}
	errorGorm = data.Database.Model(entity.Permission{}).
		Where("CompanyId = ? AND UserId = ? AND RoleId = ?", companyId, userId, roleId).
		First(&entityPermission).
		Error

	if errorGorm != gorm.ErrRecordNotFound {
		http.Error(w, "Permissions already set", http.StatusBadRequest)
		return
	}

	entityPermission.CompanyID = companyId
	entityPermission.UserID = userId
	entityPermission.RoleID = roleId

	errorGorm = data.Database.Create((&entityPermission)).Error

	if errorGorm == gorm.ErrRecordNotFound {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (data *HandlerData) RemovePermission(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIdString := vars["userId"]
	roleIdString := vars["roleId"]

	companyIdString, _ := r.Context().Value("companyId").(string)

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
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if errorGorm != nil {
		http.Error(w, errorGorm.Error(), http.StatusInternalServerError)
		return
	}

	data.Database.Delete(&entityPermission)

	w.WriteHeader(http.StatusNoContent)
}
