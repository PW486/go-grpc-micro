package main

import (
	"net/http"

	"github.com/PW486/gost/db"
	"github.com/PW486/gost/model"
	"github.com/gin-gonic/gin"
)

func getHandler(c *gin.Context) {
	var articles []model.Article
	db.Service().Find(&articles)

	c.JSON(200, gin.H{"data": articles})
}

func postHandler(c *gin.Context) {
	var newArticle model.Article
	if err := c.ShouldBindJSON(&newArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Service().Create(&newArticle)

	c.JSON(201, gin.H{"data": newArticle})
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
