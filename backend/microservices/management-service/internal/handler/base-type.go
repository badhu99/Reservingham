package handler

import "gorm.io/gorm"

type HandlerData struct {
	// database
	Database *gorm.DB
}
