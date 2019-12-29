package main

import (
	"context"

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

	client := greetpb.NewGreetingServiceClient(connection)
	response, err := client.Greet(
		context.Background(),
		&greetpb.GreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Adithya Visnu",
				LastName:  "Prasetyo Putra",
			},
		},
	)
	if err != nil {
		println(err)
	}

	println("Response: ", response.GetMessage())
}
