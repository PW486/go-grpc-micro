package account

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PW486/gost/database"
	"github.com/PW486/gost/entity"
	pb "github.com/PW486/gost/protobuf/match"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

func GetHandler(c *gin.Context) {
	var accounts []entity.Account
	database.GetDB().Find(&accounts)

	c.JSON(200, gin.H{"data": accounts})
}

func PostHandler(c *gin.Context) {
	var createAccountDTO CreateAccountDTO
	if err := c.ShouldBindJSON(&createAccountDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var newAccount entity.Account
	newAccount.ID, _ = uuid.NewUUID()
	newAccount.Email = createAccountDTO.Email
	newAccount.Name = createAccountDTO.Name
	newAccount.Password, _ = bcrypt.GenerateFromPassword([]byte(createAccountDTO.Password), 10)

	database.GetDB().Create(&newAccount)

	c.JSON(201, gin.H{"data": newAccount})
}

func LogInHandler(c *gin.Context) {
	var logInDTO LogInDTO
	if err := c.ShouldBindJSON(&logInDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var account entity.Account
	database.GetDB().Where("Email = ?", logInDTO.Email).First(&account)

	if err := bcrypt.CompareHashAndPassword(account.Password, []byte(logInDTO.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)

	c.JSON(200, gin.H{"token": ss})
}

func GetMatchHandler(c *gin.Context) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMatchClient(conn)

	match := c.Param("match")

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
}
