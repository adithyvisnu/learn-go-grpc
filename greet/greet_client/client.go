package main

import (
	"github.com/adithyvisnu/learn-go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	println("Hello World! This is pb client")

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		println(err)
	}

	defer connection.Close()

	client := greetpb.NewGreetingClient(connection)
	println(client)
}
