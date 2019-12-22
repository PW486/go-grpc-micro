package server

import (
	"context"

	"github.com/PW486/go-grpc-micro/database"
	"github.com/PW486/go-grpc-micro/protobuf/match"
	"github.com/jinzhu/gorm"
)

type matchServer struct{}

func (s *matchServer) GetMatchAccountByID(ctx context.Context, in *match.GetMatchAccountByIDRequest) (*match.GetMatchAccountByIDResponse, error) {
	var account match.GetMatchAccountByIDResponse
	if err := database.GetDB().Table("accounts").Where("id = ?", in.Id).Scan(&account).Error; gorm.IsRecordNotFoundError(err) {
		return nil, err
	}

	return &account, nil
}
