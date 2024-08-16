# Basic Go gRPC Server and Client

This project demonstrates an implementation of a gRPC server and client in Go. It includes examples of:

- Simple Unary RPC
- Server-side streaming RPC
- Client-side streaming RPC
- Bidirectional streaming RPC


# Setting up a gRPC-Go project

1. Create a new directory for your project and cd into it

   - mkdir basic-go-grpc
   -  cd basic-go-grpc
   - mkdir client server proto

3. Installing the gRPC Go plugin

   - go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
   - go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
   - export PATH="$PATH:$(go env GOPATH)/bin"

4. Create the proto file with the required services and messages in the proto directory

5. Generate .pb.go files from the proto file

    depending on what path you mention in your greet.proto file, you will either run this -

   -  protoc --go_out=. --go-grpc_out=. proto/greet.proto

    or this
   -  protoc --go_out=. --go_opt=module=github.com/akhil/basic-go-grpc --go-grpc_out=. --go-grpc_opt=module=githu
      b.com/akhil/basic-go-grpc proto/greet.proto


6. Create the server and client directories and create the main.go files with necessary controllers and services

# Running the application

1.  Install the dependencies
     go mod tidy

3.  Run the Server
     go run server/main.go

5.  Run the Client
     go run client/main.go
