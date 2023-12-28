package dto

import mssql "github.com/microsoft/go-mssqldb"

type RoleResponse struct {
	ID   mssql.UniqueIdentifier `json:"Id"`
	Name string                 `json:"Name"`
}
