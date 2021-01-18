package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/renatospaka/fc2.0-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// AddUser(client)
	// AddUserVerbose(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:                   "0",
		Name:                 "João Melão",
		Email:                "melao@joao.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:                   "0",
		Name:                 "João Melão",
		Email:                "melao@joao.com",
	}

	res, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := res.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive message: %v", err)
		}

		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id: "w1", Name: "Wesley", Email: "wes@wes.com",
		},
		&pb.User{
			Id: "r1", Name: "Renato", Email: "renato@email.com",
		},
		&pb.User{
			Id: "ro1", Name: "Rochinha", Email: "rochinha@sorvetes.com.br",
		},
		&pb.User{
			Id: "w2", Name: "Wesley 2", Email: "wes2@wes.com",
		},
		&pb.User{
			Id: "r2", Name: "Renato 2", Email: "renato2@email.com",
		},
		&pb.User{
			Id: "ro2", Name: "Rochinha 2", Email: "rochinha2@sorvetes.com.br",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 2)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}