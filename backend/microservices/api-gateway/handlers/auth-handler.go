package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/badhu99/api-gateway/models"
)

func SingIn(w http.ResponseWriter, r *http.Request) {
	log.Println("First one was pinged")

	requestUrl := "http://localhost:8080/auth/signin"
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Println("Error1: ", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error1: ", err)
		return
	}

	fmt.Printf("Response code status: %d", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Error: %s", err)
	}

	log.Printf("This is response: %s ", resBody)

	var user models.Login
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		log.Printf("Error: %s", err)
	}

	log.Printf("Username: %s, Password: %s", user.Username, user.Password)

	w.Write([]byte("Dejla dejla"))
	w.WriteHeader(http.StatusAccepted)
}
