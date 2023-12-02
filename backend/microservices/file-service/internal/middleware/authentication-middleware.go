package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/badhu99/file-service/internal/services"
)

func AuthHandler(roles []services.Role, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) != 2 {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		token := strings.TrimSpace(splitToken[1])

		claims, err := services.ValidateJwt(token, roles)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "userId", claims.Id))

		if claims.CompanyId != "" {
			r = r.WithContext(context.WithValue(r.Context(), "companyId", claims.CompanyId))
		}
		next(w, r)
	}
}

func AuthSubRouter(roles []services.Role) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer")
			if len(splitToken) != 2 {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			token := strings.TrimSpace(splitToken[1])

			claims, err := services.ValidateJwt(token, roles)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), "userId", claims.Id))

			if claims.CompanyId != "" {
				r = r.WithContext(context.WithValue(r.Context(), "companyId", claims.CompanyId))
			}

			// vars := mux.Vars(r)
			// fileName := vars["fileName"]
			filePath := r.URL.Query().Get("filePath")

			if checkFolderNameAccess(filePath, claims.CompanyId) == false {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func checkFolderNameAccess(folderPath, companyId string) bool {
	switch folderPath {
	case "temp":
		return true
	case "":
		return true
	case "test":
		return true
	case companyId:
		return true
	default:
		return false
	}
}
