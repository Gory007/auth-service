package main

import (
	"context"
	"log"
	"net"

	"proto"

	"google.golang.org/grpc"
)

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

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterAuthServiceServer(s, &server{})
	log.Println("gRPC server started on :50051")
	s.Serve(lis)
}
