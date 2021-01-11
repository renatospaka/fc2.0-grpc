package main

import (
	"log"
	"net"

	"github.com/renatospaka/fc2.0-grpc/pb"
	"github.com/renatospaka/fc2.0-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)
	
	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}