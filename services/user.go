package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/renatospaka/fc2.0-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// 	AddUser(context.Context, *User) (*User, error)
func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	fmt.Println(req.Name)
	return &pb.User{
		Id:    "1234",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

// 	AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
func (*UserService) AddUserVerbose(req *pb.User, res pb.UserService_AddUserVerboseServer) error {
	res.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	res.Send(&pb.UserResultStream{
		Status: "Creating",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	res.Send(&pb.UserResultStream{
		Status: "User has been created",
		User: &pb.User{
			Id:    "1234",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)

	res.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "1234",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)

	return nil
}

// AddUsers(ctx context.Context, opts ...grpc.CallOption) (UserService_AddUsersClient, error)
func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}

		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})
		fmt.Println("Adding ", req.GetName())
	}
}
