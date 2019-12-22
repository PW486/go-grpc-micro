package router

import (
	"github.com/PW486/gost/service/account"
	"github.com/gin-gonic/gin"
)

// Init is to return the gin router.
func Init() *gin.Engine {
	r := gin.Default()

	r.GET("/accounts", account.GetAccountsHandler)
	r.GET("/accounts/:id", account.GetAccountByIDHandler)
	r.POST("/accounts", account.PostAccountHandler)
	r.DELETE("/accounts/:id", account.DeleteAccountHandler)

	r.POST("/login", account.LogInHandler)

	return r
}
