package main

import (
	"context"

	pb "github.com/someshnayak29/gRPC-Bidirectional-streaming/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
