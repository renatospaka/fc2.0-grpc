package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %w", err)
	}

	grpcServer := grpc.NewServer()
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}