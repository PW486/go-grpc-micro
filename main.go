package main

import (
	"fmt"

	"github.com/PW486/gost/config"
	"github.com/PW486/gost/database"
	"github.com/PW486/gost/entity"
	"github.com/PW486/gost/router"
	"github.com/PW486/gost/server"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	config.Init()

	db := database.Init()
	db.AutoMigrate(&entity.Account{})

	go server.Run()

	gin.SetMode(config.AppSetting.RunMode)

	router := router.Init()
	port := fmt.Sprintf(":%d", config.AppSetting.HTTPPort)

	router.Run(port)
}
