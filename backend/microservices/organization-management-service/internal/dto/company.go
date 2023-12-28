package dto

import mssql "github.com/microsoft/go-mssqldb"

type CompanyData struct {
	Name  string `json:"Name" validate:"required"`
	Name1 string `json:"Name1,omitempty"`
}

type Company struct {
	ID mssql.UniqueIdentifier `json:"Id"`
	CompanyData
}
