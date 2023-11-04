package dto

type CompanyData struct {
	Name string `json:"Name" validate:"required"`
}

type Company struct {
	ID string `json:"Id"`
	CompanyData
}

type PaginationCompany struct {
	Count int       `json:"Count"`
	Page  int       `json:"PageNumber"`
	Size  int       `json:"PageSize"`
	Items []Company `json:"Items"`
}
