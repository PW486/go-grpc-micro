package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

// Init opens and returns the database.
func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "../test.db")
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
