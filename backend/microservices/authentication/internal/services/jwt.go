package services

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/badhu99/authentication/internal/entity"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type JwtData struct {
	jwt.StandardClaims
	Id       string
	Username string
	Roles    []string
}

func GenerateJwt(user entity.User) (string, error) {

	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatalln("Check .env file: ", err)
	}

	jwtKey := os.Getenv("JWT_SECRET")

	claims := JwtData{
		StandardClaims: jwt.StandardClaims{Audience: "Audience", ExpiresAt: time.Now().Add(12 * time.Hour).Unix(), IssuedAt: time.Now().Unix(), Issuer: "Issuer", Subject: "Subject"},
		Id:             user.ID.String(),
		Username:       user.Username,
		// Role:           user.Roles,
	}

	var rolesString []string
	for _, role := range user.Roles {
		rolesString = append(rolesString, role.Name)
	}

	claims.Roles = rolesString

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenString.SignedString([]byte(jwtKey))
}

func ValidateJwt(tokenString string) (string, []string, error) {

	jwtKey := os.Getenv("JWT_SECRET")

	token, err := jwt.ParseWithClaims(tokenString, &JwtData{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		return "", []string{}, err
	}

	claims := &JwtData{}

	claims, ok := token.Claims.(*JwtData)
	if !ok {
		return "", []string{}, errors.New("Could not parse claims!")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return "", []string{}, errors.New("Token has expired!")
	}

	return claims.Id, claims.Roles, nil

}
