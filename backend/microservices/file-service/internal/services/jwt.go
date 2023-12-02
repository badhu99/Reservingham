package services

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
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

const audienceCheckValue = "file-service"

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
		return nil, fmt.Errorf("could not parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, fmt.Errorf("token has expired")
	}

	// Check audience
	claimsValid := func() bool {
		claimsArray := strings.Split(claims.Audience, ";")
		for _, key := range claimsArray {
			if key == audienceCheckValue {
				return true
			}
		}
		return false
	}()
	if !claimsValid {
		return nil, fmt.Errorf("claims are invalid")
	}

	// Check roles
	for _, inputRole := range roles {
		for _, userRole := range claims.Roles {
			if inputRole == userRole {
				return claims, nil
			}
		}
	}

	return nil, errors.New("Role not found")
}

func GenerateJwt(claims JwtData) (string, error) {

	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatalln("Check .env file: ", err)
	}

	jwtKey := os.Getenv("JWT_SECRET")

	// _ := JwtData{
	// 	StandardClaims: jwt.StandardClaims{Audience: "Audience", ExpiresAt: time.Now().Add(1 * time.Minute).Unix(), IssuedAt: time.Now().Unix(), Issuer: "Issuer", Subject: "Subject"},
	// 	Id:             "",
	// 	Username:       "",
	// }

	// var rolesString []string
	// for _, role := range user.Roles {
	// 	rolesString = append(rolesString, role.Name)
	// }
	// claims.Roles = rolesString

	// if user.Company.ID != [16]byte{} {
	// 	claims.CompanyId = user.Company.ID.String()
	// }

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenString.SignedString([]byte(jwtKey))
}
