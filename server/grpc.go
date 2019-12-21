package server

import (
	"log"
	"net"

	"github.com/PW486/gost/protobuf/match"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run is to serve gRPC server.
func Run() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	match.RegisterMatchServer(s, &matchServer{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
