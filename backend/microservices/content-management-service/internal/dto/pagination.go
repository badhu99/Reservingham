package dto

type Pagination[T any] struct {
	Count int `json:"Count"`
	Page  int `json:"PageNumber"`
	Size  int `json:"PageSize"`
	Items []T `json:"Items"`
}
