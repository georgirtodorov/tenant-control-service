package gprcserver

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// GRPCServer represents a gRPC server
type GRPCServer struct {
	server *grpc.Server
	port   string
}

// New creates a new GRPCServer
func New(port string) *GRPCServer {
	return &GRPCServer{
		server: grpc.NewServer(),
		port:   port,
	}
}

// Start starts the gRPC server
func (s *GRPCServer) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	log.Printf("gRPC server is listening on port %s", s.port)

	return s.server.Serve(listener)
}

// RegisterService registers a service with the gRPC server
func (s *GRPCServer) RegisterService(registerFunc func(*grpc.Server)) {
	registerFunc(s.server)
}

// GetServer returns the underlying grpc.Server instance.
func (s *GRPCServer) GetServer() *grpc.Server {
	return s.server
}
