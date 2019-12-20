package router

import (
	"github.com/PW486/gost/service/account"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.GET("/", account.GetHandler)
	r.POST("/", account.PostHandler)
	r.POST("/login", account.LogInHandler)
	r.GET("/:match", account.GetMatchHandler)

	return r
}
