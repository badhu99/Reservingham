package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtData struct {
	jwt.StandardClaims
	Id        string
	Username  string
	Roles     []Role
	CompanyId string `json:"CompanyId,omitempty"`
}
type Role string

const (
	User    Role = "User"
	Editor  Role = "Editor"
	Manager Role = "Manager"
	Admin   Role = "Admin"
)

func ValidateJwt(tokenString string, roles []Role) (*JwtData, error) {

	jwtKey := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, &JwtData{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims := &JwtData{}

	claims, ok := token.Claims.(*JwtData)
	if !ok {
		return nil, errors.New("Could not parse claims!")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("Token has expired!")
	}

	for _, inputRole := range roles {
		for _, userRole := range claims.Roles {
			if inputRole == userRole {
				return claims, nil
			}
		}
	}

	return nil, errors.New("Role not found")
}
