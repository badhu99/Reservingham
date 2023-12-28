package services

import (
	"github.com/badhu99/organization-management-service/internal/dto"
	"github.com/badhu99/organization-management-service/internal/entity"
)

func (data *BaseServiceData) GetRoles(pageNumber, pageSize int) dto.Pagination[dto.RoleResponse] {
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

	return dto.Pagination[dto.RoleResponse]{
		Count: int(count),
		Page:  pageNumber,
		Size:  pageSize,
		Items: responseRoles,
	}
}
