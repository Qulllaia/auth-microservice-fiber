package gRPC

import (
	context "context"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type server struct {
	UnimplementedUserServiceServer
}

func InitGRPC() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterUserServiceServer(s, &server{})

	go func() {
		log.Println("gRPC server running on :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}

func (s *server) GetUser(ctx context.Context, req *UserRequest) (*UserResponse, error) {
	
	return &UserResponse{
		Id:    req.UserId,
		Email: "john@example.com",
	}, nil
}