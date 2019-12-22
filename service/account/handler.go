package account

import (
	"log"
	"net/http"
	"time"

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

	var matchAccount *match.GetMatchAccountByIDResponse
	if account.MatchID != nil {
		matchID := account.MatchID.String()
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

	account, err := createAccount(payload)
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

	err := removeAccountByID(id)
	if err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

// LogInHandler verifies the account and returns token.
func LogInHandler(c *gin.Context) {
	var payload LogInDTO
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := findAccountByEmail(payload.Email)
	if account == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email or Password Not Valid"})
		return
	}
	if err := bcrypt.CompareHashAndPassword(account.Password, []byte(payload.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email or Password Not Valid"})
		return
	}

	var expiredTime int64 = 60 * 60 * 24
	claims := &jwt.StandardClaims{ExpiresAt: time.Now().Unix() + expiredTime}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signingKey := []byte("PW486")
	signedString, err := token.SignedString(signingKey)
	if err != nil {
		log.Print(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": signedString})
}
