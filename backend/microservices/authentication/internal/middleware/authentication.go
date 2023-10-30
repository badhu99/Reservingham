package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/badhu99/authentication/internal/services"
)

func AuthenticateWrapper(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		token := strings.TrimSpace(splitToken[1])
		id, _, err := services.ValidateJwt(token)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		log.Println(id)
		r = r.WithContext(context.WithValue(r.Context(), "userId", id))
		next(w, r)
	}
}
