package db

import (
	"github.com/PW486/gost/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func DBopen() {
	var err error
	DB, err = gorm.Open("sqlite3", "../test.db")
	if err != nil {
		panic("failed to connect database")
	}
}

func DBmigraion() {
	if DB != nil {
		DB.AutoMigrate(&model.Article{})
	}
}
