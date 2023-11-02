package dto

type Login struct {
	Username string `json:"Username"    validate:"required"`
	Password string `json:"Password"    validate:"required"`
}
