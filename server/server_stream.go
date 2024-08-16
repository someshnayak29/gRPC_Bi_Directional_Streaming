package main

import (
	"log"
	"time"

	pb "github.com/someshnayak29/gRPC-Bidirectional-streaming/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {

	log.Printf("Got request with names %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}

		// send response as stream
		if err := stream.Send(res); err != nil {
			return err
		}

		// 2 second delay to simulate a long running process
		time.Sleep(time.Second * 2)
	}
	return nil
}
