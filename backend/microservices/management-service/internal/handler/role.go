package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/badhu99/management-service/internal/dto"
	"github.com/badhu99/management-service/internal/entity"
)

func (data *HandlerData) GetRoles(w http.ResponseWriter, r *http.Request) {
	pageNumber, err := strconv.Atoi(r.URL.Query().Get("pageNumber"))
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		pageSize = 12
	}

	var count int64
	entityRole := []entity.Role{}

	data.Database.Model([]entity.Role{}).
		Count(&count).
		Offset((pageNumber - 1) * pageSize).Limit(pageSize).
		Where("Level < 4").
		Find(&entityRole)

	responseRoles := []dto.RoleResponse{}
	for _, r := range entityRole {
		responseRoles = append(responseRoles, dto.RoleResponse{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	responseData := dto.Pagination[dto.RoleResponse]{
		Count: int(count),
		Page:  pageNumber,
		Size:  pageSize,
		Items: responseRoles,
	}

	response, _ := json.Marshal(responseData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
