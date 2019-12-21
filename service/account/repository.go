package account

import (
	"log"

	"github.com/PW486/gost/protobuf/match"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

// GetAccount takes another service account.
func GetAccount(c *gin.Context, id string) *match.Account {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := match.NewMatchClient(conn)

	req := &match.GetAccountRequest{Id: id}
	res, err := client.GetAccount(c, req)
	if err != nil {
		return nil
	}

	return res
}
