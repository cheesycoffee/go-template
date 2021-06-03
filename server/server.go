package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	appConfig "github.com/cheesycoffee/go-template/config"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type (
	server struct {
		config     appConfig.Config
		httpEngine []*http.Server
		grpcEngine *grpc.Server
	}

	// Server runner
	Server interface {
		Serve()
	}
)

// New server runner
func New(c appConfig.Config) Server {
	return &server{
		config: c,
	}
}

func (s *server) Serve() {
	netListener, err := net.Listen("tcp", fmt.Sprintf(":%v", s.config.Env().ServerPort))
	if err != nil {
		panic(err)
	}

	m := cmux.New(netListener)

	if s.config.Env().UseRest {
		httpListener := m.Match(cmux.HTTP1Fast())
		httpServer := &http.Server{}
		s.httpEngine = append(s.httpEngine, httpServer)
		go httpServer.Serve(httpListener)
	}

	if s.config.Env().UseGraphQL {
		gqlListener := m.Match(cmux.HTTP1Fast())
		gqlServer := &http.Server{}
		s.httpEngine = append(s.httpEngine, gqlServer)
		go gqlServer.Serve(gqlListener)
	}

	if s.config.Env().UseGRPC {
		grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
		grpcServer := grpc.NewServer()
		s.grpcEngine = grpcServer
		go grpcServer.Serve(grpcListener)
	}

	m.Serve()
}

func (s *server) Shutdown() {
	for _, engine := range s.httpEngine {
		if engine != nil {
			engine.Shutdown(context.Background())
		}
	}

	if s.grpcEngine != nil {
		s.grpcEngine.GracefulStop()
	}
}
