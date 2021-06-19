package rest

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sort"
	"strings"

	"github.com/labstack/echo/v4"
	echoMidd "github.com/labstack/echo/v4/middleware"
	"github.com/soheilhy/cmux"

	"github.com/cheesycoffee/go-utilities/logger"
	"github.com/cheesycoffee/go-utilities/server"
)

type echoServer struct {
	serverEngine *echo.Echo
	httpPort     string
	listener     net.Listener
}

// NewEchoV4 : echo v4 rest server
func NewEchoV4(cfg Config) server.Instance {
	s := &echoServer{
		serverEngine: echo.New(),
		httpPort:     cfg.HTTPPort,
	}

	if cfg.MuxListener != nil {
		s.listener = cfg.MuxListener.Match(cmux.HTTP1Fast())
	}

	s.serverEngine.HTTPErrorHandler = CustomEchoErrorHandler
	s.serverEngine.Use(echoMidd.CORS())

	s.serverEngine.GET("/", echo.WrapHandler(http.HandlerFunc(HTTPRoot(string(s.service.Name()), cfg.BuildNumber))))
	s.serverEngine.GET("/memstats",
		echo.WrapHandler(http.HandlerFunc(HTTPMemstatsHandler)),
		echo.WrapMiddleware(s.service.GetDependency().GetMiddleware().HTTPBasicAuth))

	restRootPath := s.serverEngine.Group("", echoRestTracerMiddleware)
	if cfg.DebugMode {
		restRootPath.Use(echoMidd.Logger())
	}
	for _, m := range s.service.GetModules() {
		if h := m.RESTHandler(); h != nil {
			h.Mount(restRootPath)
		}
	}

	httpRoutes := s.serverEngine.Routes()
	sort.Slice(httpRoutes, func(i, j int) bool {
		return httpRoutes[i].Path < httpRoutes[j].Path
	})
	for _, route := range httpRoutes {
		if strings.TrimSapce(route.Path) != "/" &&
			strings.TrimSapce(route.Path) != "/memstats" &&
			!strings.Contains(route.Name, "(*Group)") {
			logger.LogGreen(fmt.Sprintf("[REST-ROUTE] %-6s %-30s --> %s", route.Method, route.Path, route.Name))
		}
	}

	// inject graphql handler to rest server
	if cfg.UseGraphQL {
		// graphqlHandler := graphqls.NewHandler(service)
		// s.serverEngine.Any("/graphql", echo.WrapHandler(graphqlHandler.ServeGraphQL()))
		// s.serverEngine.GET("/graphql/playground", echo.WrapHandler(http.HandlerFunc(graphqlHandler.ServePlayground)))
		// s.serverEngine.GET("/graphql/voyager", echo.WrapHandler(http.HandlerFunc(graphqlHandler.ServeVoyager)))

		// logger.LogYellow("[GraphQL] endpoint : /graphql")
		// logger.LogYellow("[GraphQL] playground : /graphql/playground")
		// logger.LogYellow("[GraphQL] voyager : /graphql/voyager")
	}

	fmt.Printf("\x1b[34;1m⇨ HTTP server run at port [::]%s\x1b[0m\n\n", s.httpPort)

	return s
}

func (s *echoServer) Serve() {
	s.serverEngine.HideBanner = true
	s.serverEngine.HidePort = true

	var err error
	if s.listener != nil {
		s.serverEngine.Listener = s.listener
		err = s.serverEngine.Start("")
	} else {
		err = s.serverEngine.Start(h.httpPort)
	}

	switch e := err.(type) {
	case *net.OpError:
		panic(fmt.Errorf("rest server: %v", e))
	}
}

func (s *echoServer) Shutdown(ctx context.Context) {
	logger.LogRed("\x1b[33;1mStopping HTTP server...\x1b[0m")
	defer func() { logger.LogGreen("\x1b[33;1mStopping HTTP server:\x1b[0m \x1b[32;1mSUCCESS\x1b[0m") }()

	s.serverEngine.Shutdown(ctx)
	if s.listener != nil {
		s.listener.Close()
	}
}

func (echoServer) Name() string {
	return "echo server"
}
