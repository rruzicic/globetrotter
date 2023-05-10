package gapi

import (
	"context"
	"log"
	"net"

	"github.com/rruzicic/globetrotter/bnb/account-service/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

// implement the methods from proto file here
func (s *server) CreateUser(ctx context.Context, in *pb.UserSignUpRequest) (*pb.UserResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *server) GetUserById(ctx context.Context, in *pb.UserRequestId) (*pb.UserResponse, error) {
	// TODO: implement
	return &pb.UserResponse{
		User: &pb.User{Id: in.Id, CreatedOn: timestamppb.Now(), FirstName: "pera"},
	}, nil
}

func (s *server) GetUserByEmail(ctx context.Context, in *pb.UserRequestEmail) (*pb.UserResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *server) GetUsers(pagination *pb.UserRequestPagination, stream pb.UserService_GetUsersServer) error {
	//stream.Send() // for sending individual Users via stream
	// TODO: implement
	return nil
}

func (s *server) UpdateUser(ctx context.Context, in *pb.UserUpdateRequest) (*pb.UserResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *server) DeleteUser(ctx context.Context, in *pb.UserRequestId) (*pb.UserResponse, error) {
	// TODO: implement
	return nil, nil
}

func InitServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("gRPC failed to listen")
		return
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})
	log.Printf("gRPC server listening at %v", listen.Addr())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("gRPC failed to serve: %v", err)
	}
}
