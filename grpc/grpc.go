package main

import (
	"context"
	"log"
	"net"

	"github.com/PW486/gost/db"
	pb "github.com/PW486/gost/match"
	"github.com/PW486/gost/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) GetAccount(ctx context.Context, in *pb.GetAccountRequest) (*pb.Account, error) {
	var account model.Account
	db.Service().Where("ID = ?", in.Id).First(&account)

	return &pb.Account{Id: in.Id, Email: account.Email, Name: account.Name}, nil
}
func main() {
	db.Open()
	db.Migration()

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
