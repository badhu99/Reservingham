package dto

type FileData struct {
	FileName string
	FilePath string `json:",omitempty"`
}
