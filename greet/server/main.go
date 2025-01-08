package main

import (
	"log"
	"net"

	protoBuffer "github.com/rcafalchio/go-with-grpc/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	protoBuffer.GreetServiceServer
}

var address string = "0.0.0.0:50010"

func main() {
	// Start the server
	list, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on %s", address)

	server := grpc.NewServer()
	protoBuffer.RegisterGreetServiceServer(server, &Server{})

	if err := server.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
