package server

import "context"

// Instance server
type Instance interface {
	Serve()
	Shutdown(ctx context.Context)
	Name() string
}
