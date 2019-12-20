package main

import (
	"github.com/PW486/gost/database"
	"github.com/PW486/gost/entity"
	"github.com/PW486/gost/router"
)

func main() {
	db := database.Init()
	db.AutoMigrate(&entity.Account{})

	r := router.Init()
	r.Run(":8080")
}
