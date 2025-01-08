package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	protoBuffer "github.com/rcafalchio/go-with-grpc/greet/proto"
)

var address string = "0.0.0.0:50010"

func main() {

	log.Print("Initializing gRCP client...")
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("Failed to dial server: %v", err)
		log.Fatalf("Failed to dial server: %v", err)

	}

	defer conn.Close()
	client := protoBuffer.NewGreetServiceClient(conn)
	doGreet("Ricardo", client)
	log.Print("Finishing gRCP client...")
}
