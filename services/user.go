package services

import (
	"context"
	"fmt"
	"time"

	"github.com/renatospaka/fc2.0-grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }
func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	fmt.Println(req.Name)
	return &pb.User{
		Id:    "1234",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

// type UserServiceClient interface {
// 	AddUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
// 	AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)
// }
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
