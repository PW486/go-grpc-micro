package database

import (
	"log"

	"github.com/PW486/go-grpc-micro/config"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

// Init opens and returns the database.
func Init() *gorm.DB {
	db, err := gorm.Open(config.DatabaseSetting.Type, config.DatabaseSetting.Host)
	if err != nil {
		log.Fatal(err)
	}

	database = db
	return database
}

// GetDB only returns the database.
func GetDB() *gorm.DB {
	return database
}
