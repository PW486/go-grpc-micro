package server

import (
	"context"
	"log"
	"net"

	"github.com/PW486/gost/database"
	"github.com/PW486/gost/entity"
	pb "github.com/PW486/gost/protobuf/match"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {
	var account entity.Account
	database.GetDB().Where("ID = ?", in.Id).First(&account)

	return &pb.Account{Id: in.Id, Email: account.Email, Name: account.Name}, nil
}
func Serve() {
	db := database.Init()
	db.AutoMigrate(&entity.Account{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMatchServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
