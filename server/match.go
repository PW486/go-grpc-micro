package server

import (
	"context"

	"github.com/PW486/gost/database"
	"github.com/PW486/gost/entity"
	"github.com/PW486/gost/protobuf/match"
)

type matchServer struct{}

func (s *matchServer) GetAccount(ctx context.Context, in *match.GetAccountRequest) (*match.Account, error) {
	var account entity.Account
	database.GetDB().Where("ID = ?", in.Id).First(&account)

	return &match.Account{Id: in.Id, Email: account.Email, Name: account.Name}, nil
}
