package grpc

import (
	"log"
	"net"

	pb "expense-management-system/internal/api/grpc/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExpenseServiceServer
}

// NewGRPCServer creates a new gRPC server
func NewGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterExpenseServiceServer(s, &server{})
	return s
}

// Start starts the gRPC server
func Start(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := NewGRPCServer()
	log.Printf("gRPC server listening at %v", address)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
