package dto

type DraftCreate struct {
	Name string
	DraftHistory
}

type Draft struct {
	Id   string
	Name string
}
