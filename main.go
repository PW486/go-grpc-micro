package main

import (
	"net/http"

	"github.com/PW486/gost/db"
	"github.com/PW486/gost/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getHandler(c *gin.Context) {
	var accounts []model.Account
	db.Service().Find(&accounts)

	c.JSON(200, gin.H{"data": accounts})
}

func postHandler(c *gin.Context) {
	var newAccount model.Account
	if err := c.ShouldBindJSON(&newAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newAccount.ID, _ = uuid.NewUUID()

	db.Service().Create(&newAccount)

	c.JSON(201, gin.H{"data": newAccount})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", getHandler)
	r.POST("/", postHandler)

	return r
}

func main() {
	db.Open()
	db.Migration()

	r := setupRouter()
	r.Run(":8080")
}
