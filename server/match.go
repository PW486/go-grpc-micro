package server

import (
	"context"

	"github.com/PW486/gost/database"
	"github.com/PW486/gost/protobuf/match"
	"github.com/jinzhu/gorm"
)

type matchServer struct{}

func (s *matchServer) GetAccount(ctx context.Context, in *match.GetAccountRequest) (*match.Account, error) {
	var account match.Account
	if err := database.GetDB().Table("accounts").Where("id = ?", in.Id).Scan(&account).Error; gorm.IsRecordNotFoundError(err) {
		return nil, err
	}

	return &account, nil
}
