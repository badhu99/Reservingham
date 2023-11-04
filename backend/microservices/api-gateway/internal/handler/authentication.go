package handler

import (
	"fmt"
	"net/http"
)

// @Summary		User authentication
// @Tags			Auth
// @Accept			json
// @Produce		json
// @Param login	body	dto.Login	true "Body"
// @Success		200		{object}	dto.UserResponse
// @Failure		400		{string}	string
// @Failure		401	    {object}	string
// @Router			/api/auth/signin [post]
func (data *HandlerData) SignIn(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("%s/signin", data.UrlAuth)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodPost, r.Body)
	functionHandler(w, r)
}
