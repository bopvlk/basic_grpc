package main

import (
	"log"
	"net"

	pb "github.com/bopvlk/basic_grpc/pb/test"
	"github.com/bopvlk/basic_grpc/pkg"
	"google.golang.org/grpc"
)

func main() {
	s := grpc.NewServer()
	srv := &pkg.GRPCServer{}

	pb.RegisterApiServer(s, srv)

	lis, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatal("problem with net.Listen(...) err: ", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatal("fatal error in s.Serve(lis) err: ", err)
	}

}
