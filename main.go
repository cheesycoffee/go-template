package main

import (
	"github.com/cheesycoffee/go-template/config"
	"github.com/cheesycoffee/go-template/server"
)

func main() {
	appConfig := config.New()

	appServer := server.New(appConfig)
	appServer.Serve()
}
