package dto

type RoleResponse struct {
	ID   string `json:"Id"`
	Name string `json:"Name"`
}

type PaginationRole struct {
	Count int            `json:"Count"`
	Page  int            `json:"PageNumber"`
	Size  int            `json:"PageSize"`
	Items []RoleResponse `json:"Items"`
}
