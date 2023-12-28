package handler

import (
	servicesBusinessLogic "github.com/badhu99/organization-management-service/internal/services/data-business"
)

type HandlerData struct {
	Services *servicesBusinessLogic.BaseServiceData
}
