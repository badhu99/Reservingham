package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/badhu99/content-management-service/internal/services"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type HandlerData struct {
	// database
	Database                 *gorm.DB
	UrlServiceUserManagement string
}

func HttpRequestBroker[T any](requestUrl, method string, body io.ReadCloser) (*T, int, error) {

	request, _ := http.NewRequest(method, requestUrl, body)
	request.Header.Add("Content-Type", "application/json")

	jwtClaims := services.JwtData{
		StandardClaims: jwt.StandardClaims{Audience: "Audience", ExpiresAt: time.Now().Add(1 * time.Minute).Unix(), IssuedAt: time.Now().Unix(), Issuer: "Issuer", Subject: "Subject"},
		Roles:          []services.Role{services.Admin},
	}
	jwtToken, _ := services.GenerateJwt(jwtClaims)
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", jwtToken))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer response.Body.Close()

	switch response.StatusCode {
	case http.StatusUnauthorized:
		responseBody, _ := io.ReadAll(response.Body)
		return nil, http.StatusUnauthorized, fmt.Errorf(string(responseBody))

	case http.StatusNotFound:
		responseBody, _ := io.ReadAll(response.Body)
		return nil, http.StatusNotFound, fmt.Errorf(string(responseBody))
	}

	var responseData T
	err = json.NewDecoder(response.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, http.StatusInternalServerError, err
	}
	return &responseData, 0, nil
}
