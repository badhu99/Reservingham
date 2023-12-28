package handler

import (
	"fmt"
	"net/http"
)

// @Summary		Get roles paginated.
// @Tags		Role
// @Produce		json
// @Param		pageNumber   	query     string  false  "Page number"
// @Param 		pageSize    	query     string  false  "Page size"
// @Success		200		{object}	dto.PaginationRole
// @Failure		400		{string}	string
// @Failure 	401	    {object}	string
// @Router		/api/role [get]
// @Security 	Bearer
func (data *HandlerData) GetRoles(w http.ResponseWriter, r *http.Request) {
	requestUrl := fmt.Sprintf("%s/role", data.UrlOrganizationManagement)
	functionHandler := data.HttpRequestBroker(requestUrl, http.MethodGet, r.Body)
	functionHandler(w, r)
}
