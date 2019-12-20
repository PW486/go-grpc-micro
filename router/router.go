package router

import (
	"fmt"
	"log"
	"net/http"

	pb "github.com/PW486/gost/protobuf/match"
	"github.com/PW486/gost/service/account"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Init() *gin.Engine {
	r := gin.Default()

	r.GET("/", account.GetHandler)
	r.POST("/", account.PostHandler)
	r.POST("/login", account.LogInHandler)

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMatchClient(conn)

	r.GET("/:match", func(c *gin.Context) {
		match := c.Param("match")

		// Contact the server and print out its response.
		req := &pb.GetAccountRequest{Id: match}
		res, err := client.GetAccount(c, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res),
		})
	})

	return r
}
