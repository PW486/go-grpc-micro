package account

import (
	"log"

	"github.com/PW486/gost/database"
	"github.com/PW486/gost/entity"
	"github.com/PW486/gost/protobuf/match"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

func FindAccounts() *[]FindAccountResponse {
	var accounts []FindAccountResponse
	database.GetDB().Table("accounts").Where("deleted_at is null").Scan(&accounts)

	return &accounts
}

func FindAccountByID(id string) *FindAccountResponse {
	var account FindAccountResponse
	if database.GetDB().Table("accounts").Where("id = ?", id).Scan(&account).RecordNotFound() {
		return nil
	}

	return &account
}

// FindMatchAccountByID takes another service account.
func FindMatchAccountByID(c *gin.Context, id string) *match.Account {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Print(err)
		return nil
	}
	defer conn.Close()

	client := match.NewMatchClient(conn)

	req := &match.GetAccountRequest{Id: id}
	res, err := client.GetAccount(c, req)
	if err != nil {
		log.Print(err)
		return nil
	}

	return res
}

func CreateAccount(payload CreateAccountDTO) (*entity.Account, error) {
	var newAccount entity.Account

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	newAccount.ID = uuid
	newAccount.Email = payload.Email
	newAccount.Name = payload.Name

	password, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		return nil, err
	}
	newAccount.Password = password
	newAccount.Match = payload.Match

	if err := database.GetDB().Create(&newAccount).Error; err != nil {
		return nil, err
	}

	return &newAccount, nil
}

func RemoveAccount(account *FindAccountResponse) error {
	if err := database.GetDB().Where("ID = ?", account.ID).Delete(&entity.Account{}).Error; err != nil {
		return err
	}

	return nil
}
