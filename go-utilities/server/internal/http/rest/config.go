package rest

import (
	"github.com/soheilhy/cmux"
	"github.com/cheesycoffee/go-utilities/server"
)

// Config : echo server configuration
type Config struct {
	DebugMode      bool
	BuildNumber    string
	UseGraphQL     bool
	ServerInstance server.Instance
	HTTPPort       string
	MuxListener    cmux.CMux
}
