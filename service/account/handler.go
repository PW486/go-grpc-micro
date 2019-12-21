package account

import (
	"net/http"

	"github.com/PW486/gost/database"
	"github.com/PW486/gost/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GetHandler(c *gin.Context) {
	var accounts []entity.Account
	database.GetDB().Find(&accounts)

	c.JSON(200, gin.H{"data": accounts})
}

func GetByIdHandler(c *gin.Context) {
	id := c.Param("id")

	var account entity.Account
	database.GetDB().Where("ID = ?", id).First(&account)

	matchAccount := GetAccount(c, account.Match.String())

	c.JSON(200, gin.H{"data": account, "match": matchAccount})
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
	newAccount.Match = &createAccountDTO.Match

	database.GetDB().Create(&newAccount)

	c.JSON(201, gin.H{"data": newAccount})
}

func DeleteHandler(c *gin.Context) {
	id := c.Param("id")

	var account entity.Account
	database.GetDB().Where("ID = ?", id).First(&account)

	c.JSON(200, gin.H{"data": account})
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
