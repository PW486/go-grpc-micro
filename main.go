package main

import (
	"github.com/PW486/gost/database"
	"github.com/PW486/gost/entity"
	"github.com/PW486/gost/router"
	"github.com/PW486/gost/server"
)

func main() {
	db := database.Init()
	db.AutoMigrate(&entity.Account{})

	go server.Serve()

	r := router.Init()
	r.Run(":8080")
}
