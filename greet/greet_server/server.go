package main

import (
	"net"

	"github.com/adithyvisnu/learn-go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	println("Hello World")

	// 50051 default grpc port
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		print(err)
	}

	grpcServer := grpc.NewServer()
	greetpb.RegisterGreetingServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		print(err)
	}
	println("Server TCP is listening...")
}
