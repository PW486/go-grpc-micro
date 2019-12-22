package server

import (
	"fmt"
	"log"
	"net"

	"github.com/PW486/go-grpc-micro/config"
	"github.com/PW486/go-grpc-micro/protobuf/match"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run serves gRPC server.
func Run() {
	port := fmt.Sprintf(":%d", config.AppSetting.RPCPort)

	lis, err := net.Listen("tcp", port)
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
