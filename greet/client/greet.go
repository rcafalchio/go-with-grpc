package main

import (
	"context"
	"log"

	protoBuffer "github.com/rcafalchio/go-with-grpc/greet/proto"
)

func doGreet(name string, client protoBuffer.GreetServiceClient) {
	log.Printf("Greeting ...")
	res, err := client.Greet(context.Background(), &protoBuffer.GreetRequest{
		FirstName: name,
	})

	if err != nil {
		log.Fatalf("Failed to greet: %v", err)
	}

	log.Printf("Greeting: %s", res.Result)

}
