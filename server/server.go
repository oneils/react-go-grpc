package main

import (
	"context"
	"github.com/oneils/http2-demo/hellopb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	hellopb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, request *hellopb.HelloRequest) (*hellopb.HelloResponse, error) {
	log.Printf("Received: %v", request.GetName())
	return &hellopb.HelloResponse{
		Message: "Hello " + request.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8000")

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer()
	hellopb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving: %v", err)
	}
}
