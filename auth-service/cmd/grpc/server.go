// cmd/grpc/server.go
package grpcserver

import (
	"context"
	"log"
	"net"

	"github.com/Gory007/auth-service/proto"
	"google.golang.org/grpc"
)

func Start() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterAuthServiceServer(s, &server{})
	log.Println("gRPC server started on :50051")
	s.Serve(lis)
}

type server struct {
	proto.UnimplementedAuthServiceServer
}

func (s *server) GetUserProfile(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		Name:      "John Doe",
		Email:     "john@example.com",
		AvatarUrl: "https://example.com/avatar.jpg",
	}, nil
}
