package dto

import mssql "github.com/microsoft/go-mssqldb"

type User struct {
	ID mssql.UniqueIdentifier `json:"Id"`
	UserData
}

type UserData struct {
	Email     string `json:"Email"`
	Username  string `json:"Username"`
	Password  string `json:"Password,omitempty"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
}

type UserResponse struct {
	User
	Roles []RoleResponse
}
