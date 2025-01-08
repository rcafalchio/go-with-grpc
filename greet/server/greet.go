package main

import (
	"context"

	protoBuffer "github.com/rcafalchio/go-with-grpc/greet/proto"
)

func (server *Server) Greet(ctx context.Context, req *protoBuffer.GreetRequest) (*protoBuffer.GreetResponse, error) {
	res := &protoBuffer.GreetResponse{
		Result: "Hello " + req.FirstName,
	}
	return res, nil
}
