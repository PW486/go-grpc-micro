package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var database *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("sqlite3", "../test.db")
	if err != nil {
		panic("failed to connect database.")
	}
	database = db

	return database
}

func GetDB() *gorm.DB {
	return database
}
