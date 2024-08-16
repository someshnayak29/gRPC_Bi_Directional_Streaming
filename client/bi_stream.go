package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/someshnayak29/gRPC-Bidirectional-streaming/proto"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming has Started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while Streaming %v", err)
			}

			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(time.Second * 2)
	}

	stream.CloseSend()
	<-waitc // will stop ending of process until all transactions are done
	log.Printf("Bidirectional Streaming Finished")
}
