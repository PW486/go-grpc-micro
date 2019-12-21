package router

import (
	"github.com/PW486/gost/service/account"
	"github.com/gin-gonic/gin"
)

// Init is to return the gin router.
func Init() *gin.Engine {
	r := gin.Default()

	r.GET("/accounts", account.GetHandler)
	r.GET("/accounts/:id", account.GetByIdHandler)
	r.POST("/accounts", account.PostHandler)
	r.DELETE("/accounts/:id", account.DeleteHandler)

	r.POST("/login", account.LogInHandler)

	return r
}
