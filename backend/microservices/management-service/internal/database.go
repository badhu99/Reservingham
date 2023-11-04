package server

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	logger "gorm.io/gorm/logger"
)

func Database() *gorm.DB {
	err := godotenv.Load("./internal/.env")
	if err != nil {
		log.Fatalln("Check .env file: ", err)
	}

	s := os.Getenv("SQL_DATABASE_URL")

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	}
	db, err := gorm.Open(sqlserver.Open(s), gormConfig)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
