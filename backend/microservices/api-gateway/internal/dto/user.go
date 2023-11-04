package dto

type Login struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type UserResponse struct {
	Id           string `json:"Password"`
	Username     string `json:"Username"`
	Email        string `json:"Email"`
	AccessToken  string `json:"AccessToken"`
	RefreshToken string `json:"RefreshToken"`
}

type User struct {
	ID string `json:"Id"`
	UserData
}

type UserData struct {
	Email    string `json:"Email"`
	Username string `json:"Username"`
	Password string `json:"Password,omitempty"`
}

type UserDataResponse struct {
	User
	Roles []RoleResponse
}

type PaginationUsers struct {
	Count int    `json:"Count"`
	Page  int    `json:"PageNumber"`
	Size  int    `json:"PageSize"`
	Items []User `json:"Items"`
}
