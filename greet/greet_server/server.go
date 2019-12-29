package main

import (
	"context"
	"fmt"
	"net"

	"github.com/adithyvisnu/learn-go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function is invoked with, ", req)
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	result := "Hello you, " + firstName + " " + lastName
	return &greetpb.GreetResponse{
		Message: result,
	}, nil
}

func main() {
	println("Hello World")

	// 50051 default grpc port
	listener, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		print(err)
	}

	grpcServer := grpc.NewServer()
	greetpb.RegisterGreetingServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(listener); err != nil {
		print(err)
	}
	println("Server TCP is listening...")
}
