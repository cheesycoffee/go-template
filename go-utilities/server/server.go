package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type server struct {
	service   Service
	instances []Instance
}

// Server utility
type Server interface {
	Serve()
}

// New server
func New(service Service) Server {
	return &server{
		service: service,
	}
}

// Run start app
func (s *server) Serve() {
	if len(s.instances) == 0 {
		panic("No server/worker running")
	}

	errServe := make(chan error)
	for _, instance := range s.instances {
		go func(inst Instance) {
			defer func() {
				if r := recover(); r != nil {
					errServe <- fmt.Errorf("%v", r)
				}
			}()
			inst.Serve()
		}(instance)
	}

	quitSignal := make(chan os.Signal, 1)
	signal.Notify(quitSignal, os.Interrupt, syscall.SIGTERM)

	log.Printf("Serverlication \x1b[32;1m%s\x1b[0m ready to run\n\n", s.service.Name())

	select {
	case e := <-errServe:
		panic(e)
	case <-quitSignal:
		s.shutdown(quitSignal)
	}
}

// graceful shutdown all server, panic if there is still a process running when the request exceed given timeout in context
func (s *server) shutdown(forceShutdown chan os.Signal) {
	fmt.Println("\x1b[34;1mGracefully shutdown... (press Ctrl+C again to force)\x1b[0m")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for _, server := range s.instances {
			server.Shutdown(ctx)
		}
		done <- struct{}{}
	}()

	select {
	case <-done:
		log.Println("\x1b[32;1mSuccess shutdown all server & worker\x1b[0m")
	case <-forceShutdown:
		log.Println("\x1b[31;1mForce shutdown server & worker\x1b[0m")
		cancel()
	case <-ctx.Done():
		log.Println("\x1b[31;1mContext timeout\x1b[0m")
	}
}
