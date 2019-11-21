package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Article required Title, Text
type Article struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
}

var articles []Article

func getHandler(c *gin.Context) {
	// articles = append(articles, Article{Title: "testTitle", Text: "testText"})

	c.JSON(200, gin.H{"data": articles})
}

func postHandler(c *gin.Context) {
	var newArticle Article
	if err := c.ShouldBindJSON(&newArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	articles = append(articles, newArticle)

	c.JSON(201, gin.H{"data": newArticle})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", getHandler)
	r.POST("/", postHandler)

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
