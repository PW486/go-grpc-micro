package router

import (
	"github.com/PW486/go-grpc-micro/service/account"
	"github.com/gin-gonic/gin"
)

// Init returns the gin router.
func Init() *gin.Engine {
	r := gin.Default()

	r.GET("/accounts", account.GetAccountsHandler)
	r.GET("/accounts/:id", account.GetAccountByIDHandler)
	r.POST("/accounts", account.PostAccountHandler)
	r.DELETE("/accounts/:id", account.DeleteAccountHandler)

	r.POST("/login", account.LogInHandler)

	return r
}
