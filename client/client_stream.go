package main

import (
	"context"
	"log"
	"time"

	pb "github.com/someshnayak29/gRPC-Bidirectional-streaming/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client Streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}

		log.Printf("Sent the request with name: %s", name)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming Finished")
	if err != nil {
		log.Fatalf("Error while Receiving %v", err)
	}

	log.Printf("%v", res.Messages)
}
