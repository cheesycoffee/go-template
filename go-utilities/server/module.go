package server

import (
	"github.com/labstack/echo"
	"google.golang.org/grpc"
)

// ModuleType server
type ModuleType string

// Module server
type Module interface {
	RESTHandler() RESTHandler
	GRPCHandler() GRPCHandler
	GraphQLHandler() GraphQLHandler
	WorkerHandler(workerType Worker) WorkerHandler
	Name() ModuleType
}

// RESTHandler delivery factory for REST handler
type RESTHandler interface {
	Mount(group *echo.Group)
}

// GRPCHandler delivery factory for GRPC handler
type GRPCHandler interface {
	Register(server *grpc.Server, middlewareGroup *MiddlewareGroup)
}

// GraphQLHandler delivery factory for GraphQL resolver handler
type GraphQLHandler interface {
	Query() interface{}
	Mutation() interface{}
	Subscription() interface{}
	RegisterMiddleware(group *MiddlewareGroup)
}

// WorkerHandler delivery factory for all worker handler
type WorkerHandler interface {
	MountHandlers(group *WorkerHandlerGroup)
}
