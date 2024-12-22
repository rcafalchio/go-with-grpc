package main

import (
	"log"
	"net"

	pb "github.com/rcafalchio/grpc-with-go/greet/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var address string = "0.0.0.0:5000"

func main() {
	// Start the server
	list, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Listening on %s", address)

	s := grpc.NewServer()

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
