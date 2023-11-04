package handler

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	dto "github.com/badhu99/authentication/internal/dto"
	"github.com/badhu99/authentication/internal/entity"
	"github.com/badhu99/authentication/internal/services"
	"github.com/badhu99/authentication/internal/utility"
)

func (handlerData *HandlerData) SignIn(w http.ResponseWriter, r *http.Request) {

	var userLogin dto.Login

	// Validate request
	e, statusCode := utility.ValidateBody(&userLogin, r.Body)
	if e != nil {
		http.Error(w, e.Error(), statusCode)
		return
	}

	// find user entity
	var userEntity entity.User

	handlerData.Database.
		Model(userEntity).
		Where(entity.User{Username: userLogin.Username}).
		Find(&userEntity).
		Joins("JOIN Permission ON Permission.UserId = [User].Id").
		Preload("UserRole").
		Scan(&userEntity.Permissions).
		Joins("JOIN Role ON Permission.RoleId = [Role].Id").
		Preload("Role").
		Scan(&userEntity.Roles).
		Joins("JOIN Company ON Permission.CompanyId = [Company].Id").
		Preload("Company").
		Scan(&userEntity.Company)

	// Check if passwords match
	password, _ := base64.StdEncoding.DecodeString(userEntity.PasswordHash)
	salt, _ := base64.StdEncoding.DecodeString(userEntity.Salt)

	passwordMatch := services.ValidatePassword(password, []byte(userLogin.Password), salt)
	if !passwordMatch {
		http.Error(w, "Username or password is incorect", http.StatusUnauthorized)
		return
	}

	// Generate token
	jwtToken, err := services.GenerateJwt(userEntity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userData := dto.User{
		Id:          userEntity.ID.String(),
		Username:    userEntity.Username,
		Email:       userEntity.Email,
		AccessToken: jwtToken,
	}

	responseData, err := json.Marshal(userData)
	if err != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(responseData)
}
