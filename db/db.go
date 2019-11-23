package db

import (
	"github.com/PW486/gost/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func Service() *gorm.DB {
	return db
}

func Open() {
	var err error
	db, err = gorm.Open("sqlite3", "../test.db")
	if err != nil {
		panic("failed to connect database")
	}
}

func Migration() {
	if db != nil {
		db.AutoMigrate(&model.Article{})
	}
}
