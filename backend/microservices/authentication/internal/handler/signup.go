package handler

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/badhu99/authentication/internal/dto"
	"github.com/badhu99/authentication/internal/entity"
	"github.com/badhu99/authentication/internal/services"
	"github.com/badhu99/authentication/internal/utility"
	"gorm.io/gorm"
)

func (handlerData *HandlerData) SignUp(w http.ResponseWriter, r *http.Request) {

	var userRegister dto.Register

	// Check for validaty of the body and deserialize into a struct
	e, statusCode := utility.ValidateBody(&userRegister, r.Body)
	if e != nil {
		http.Error(w, e.Error(), statusCode)
		return
	}

	// Check if password match
	if userRegister.Password != userRegister.PasswordRepeat {
		http.Error(w, "Password are not a match", http.StatusBadRequest)
	}

	// Check if user with email or username already exists
	userEntity := entity.User{
		// Username: userRegister.Username,
	}

	errorGorm := handlerData.Database.Where(
		&entity.User{Username: userRegister.Username},
	).Or(
		handlerData.Database.Where(&entity.User{Email: userRegister.Email}),
	).First(&userEntity).Error

	if errorGorm != gorm.ErrRecordNotFound {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}

	// Generates salt
	salt, e := services.GenerateRandomSalt(16)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
	}

	// Hash password
	password, e := services.HashPassword([]byte(userRegister.Password), salt)
	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)
		return
	}

	// Populate user entity data and save to database
	userEntity.Username = userRegister.Username
	userEntity.Email = userRegister.Email
	userEntity.PasswordHash = base64.StdEncoding.EncodeToString(password)
	userEntity.Salt = base64.StdEncoding.EncodeToString(salt)
	handlerData.Database.Create(&userEntity)

	// Get Role entity data
	var roleEntity entity.Role
	errorGorm = handlerData.Database.Where(&entity.Role{Name: "User"}).First(&roleEntity).Error
	if errorGorm != nil {
		log.Printf(errorGorm.Error())
		http.Error(w, errorGorm.Error(), http.StatusBadRequest)
		return
	}
	// Create UserRole entity and populate data
	userRoleEntity := entity.Permission{
		UserID: userEntity.ID,
		RoleID: roleEntity.ID,
	}
	handlerData.Database.Create(&userRoleEntity)

	// Create JWT token
	accessToken, err := services.GenerateJwt(userEntity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Remap into response struct
	userResponse := dto.User{
		Id:          userEntity.ID.String(),
		Username:    userEntity.Username,
		Email:       userEntity.Email,
		AccessToken: accessToken,
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}
