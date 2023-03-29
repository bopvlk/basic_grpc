package pkg

import (
	"context"
	"errors"

	pb "github.com/bopvlk/basic_grpc/pb/test"
)

type GRPCServer struct {
	pb.UnimplementedApiServer
}

func (*GRPCServer) GetTest(ctx context.Context, req *pb.GetRequestTest) (*pb.ResponseTest, error) {
	if req.Text == "Hello" {
		return &pb.ResponseTest{Result: "World"}, nil
	}
	return nil, errors.New("wrong request text; requesttext must be 'Hello'")
}
