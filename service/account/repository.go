package account

import (
	"log"

	"github.com/PW486/gost/config"
	"github.com/PW486/gost/database"
	"github.com/PW486/gost/entity"
	"github.com/PW486/gost/protobuf/match"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

// FindAccounts finds all accounts in the database.
func FindAccounts() *[]FindAccountResponse {
	var accounts []FindAccountResponse
	database.GetDB().Table("accounts").Where("deleted_at is null").Scan(&accounts)

	return &accounts
}

// FindAccountByID finds one account by id in the database.
func FindAccountByID(id string) *FindAccountResponse {
	var account FindAccountResponse
	if database.GetDB().Table("accounts").Where("id = ?", id).Scan(&account).RecordNotFound() {
		return nil
	}

	return &account
}

// findAccountByEmail finds one account by email in the database.
func findAccountByEmail(email string) *entity.Account {
	var account entity.Account
	if database.GetDB().Where("Email = ?", email).First(&account).RecordNotFound() {
		return nil
	}

	return &account
}

// FindMatchAccountByID takes another service's account.
func FindMatchAccountByID(c *gin.Context, id string) *match.GetMatchAccountByIDResponse {
	conn, err := grpc.Dial(config.AppSetting.MatchHost, grpc.WithInsecure())
	if err != nil {
		log.Print(err)
		return nil
	}
	defer conn.Close()

	client := match.NewMatchClient(conn)

	req := &match.GetMatchAccountByIDRequest{Id: id}
	res, err := client.GetMatchAccountByID(c, req)
	if err != nil {
		log.Print(err)
		return nil
	}

	return res
}

// createAccount creates one account for the database.
func createAccount(payload CreateAccountDTO) (*entity.Account, error) {
	var newAccount entity.Account

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	newAccount.ID = uuid
	newAccount.Email = payload.Email
	newAccount.Name = payload.Name

	password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), config.AppSetting.BcryptCost)
	if err != nil {
		return nil, err
	}
	newAccount.Password = password
	newAccount.MatchID = payload.MatchID

	if err := database.GetDB().Create(&newAccount).Error; err != nil {
		return nil, err
	}

	return &newAccount, nil
}

// removeAccountByID removes one account by id for the database.
func removeAccountByID(id string) error {
	if err := database.GetDB().Where("id = ?", id).Delete(&entity.Account{}).Error; err != nil {
		return err
	}

	return nil
}
