package grpc

import (
	"context"
	"log"
	"net"

	"github.com/soheilhy/cmux"
	"github.com/cheesycoffee/go-utilities/server"
	"google.golang.org/grpc"
)

type grpcServer struct {
	serverEngine *grpc.Server
	listener     net.Listener
}

// New GRPC server
func New(muxListener cmux.CMux) server.Server {
	return &grpcServer{}
}

func (s *grpcServer) Serve() {
	if err := s.serverEngine.Serve(s.listener); err != nil {
		log.Println("GRPC: Unexpected Error", err)
	}
}

func (s *grpcServer) Shutdown(ctx context.Context) {
	log.Println("\x1b[33;1mStopping GRPC server...\x1b[0m")
	defer func() { log.Println("\x1b[33;1mStopping GRPC server:\x1b[0m \x1b[32;1mSUCCESS\x1b[0m") }()

	s.serverEngine.GracefulStop()
	s.listener.Close()
}
