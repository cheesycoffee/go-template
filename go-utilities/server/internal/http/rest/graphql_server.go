package rest

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/soheilhy/cmux"
	"github.com/cheesycoffee/go-utilities/server"
)

type graphqlServer struct {
	httpEngine *http.Server
	listener   net.Listener
	service    server.Service
}

// NewGraphQL server
func NewGraphQL(muxListener cmux.CMux, service server.Service) server.Server {
	return &graphqlServer{}
}

func (s *graphqlServer) Serve() {
	var err error
	if s.listener != nil {
		err = s.httpEngine.Serve(s.listener)
	} else {
		err = s.httpEngine.ListenAndServe()
	}

	switch e := err.(type) {
	case *net.OpError:
		panic(fmt.Errorf("gql http server: %v", e))
	}
}

func (s *graphqlServer) Shutdown(ctx context.Context) {
	log.Println("\x1b[33;1mStopping GraphQL HTTP server...\x1b[0m")
	defer func() { log.Println("\x1b[33;1mStopping GraphQL HTTP server:\x1b[0m \x1b[32;1mSUCCESS\x1b[0m") }()

	s.httpEngine.Shutdown(ctx)
	if s.listener != nil {
		s.listener.Close()
	}
}
