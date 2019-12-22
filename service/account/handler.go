package account

import (
	"log"
	"net/http"

	"github.com/PW486/gost/database"
	"github.com/PW486/gost/entity"
	"github.com/PW486/gost/protobuf/match"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetAccountsHandler finds all accounts.
func GetAccountsHandler(c *gin.Context) {
	accounts := FindAccounts()

	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

// GetAccountByIDHandler finds one account.
func GetAccountByIDHandler(c *gin.Context) {
	id := c.Param("id")
	account := FindAccountByID(id)
	if account == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Account Not Found"})
		return
	}

	var matchAccount *match.Account
	if account.Match != nil {
		matchID := account.Match.String()
		matchAccount = FindMatchAccountByID(c, matchID)
	}

	c.JSON(http.StatusOK, gin.H{"account": account, "matchAccount": matchAccount})
}

// PostAccountHandler creates one account.
func PostAccountHandler(c *gin.Context) {
	var payload CreateAccountDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account, err := CreateAccount(payload)
	if err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"account": account})
}

// DeleteAccountHandler removes one account.
func DeleteAccountHandler(c *gin.Context) {
	id := c.Param("id")
	account := FindAccountByID(id)
	if account == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Account Not Found"})
		return
	}

	err := RemoveAccount(account)
	if err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// LogInHandler verifies the account and returns token.
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

	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)

	c.JSON(http.StatusOK, gin.H{"token": ss})
}
