package dto

import mssql "github.com/microsoft/go-mssqldb"

type CompanyData struct {
	Name string `json:"Name"`
}

type Company struct {
	ID mssql.UniqueIdentifier `json:"Id"`
	CompanyData
}
