package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/diegoholiveira/gobalancer/nodes"
	"github.com/diegoholiveira/gobalancer/nodes/container"
	"github.com/diegoholiveira/gobalancer/nodes/manager"
	"github.com/diegoholiveira/gobalancer/server"
)

const (
	gRPCPort string = ":8000"
	HTTPPort string = ":8080"
)

func main() {
	grpcListener, err := net.Listen("tcp", gRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(
		shutdownSignal,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	container := container.NewContainer()

	srv := newHTTPServer(ctx, cancel, container)
	grpc := newGRPCServer(container)

	go func() {
		log.Printf("http server at %s\n", HTTPPort)

		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server error: %v", err)
		}
	}()

	go func() {
		log.Printf("grpc server at %s\n", HTTPPort)

		if err := grpc.Serve(grpcListener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ticker := time.NewTicker(10 * time.Millisecond)

out:
	for {
		select {
		case <-shutdownSignal:
			ticker.Stop()

			break out
		case <-ticker.C:
			container.Update()
		}
	}

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := srv.Shutdown(gracefullCtx); err != nil {
		log.Printf("http server shutdown error: %v\n", err)
	} else {
		log.Printf("http server gracefully stopped\n")
	}

	grpc.GracefulStop()

	log.Printf("grpc server gracefully stopped\n")
}

func newHTTPServer(ctx context.Context, cancel context.CancelFunc, container *container.Container) *http.Server {
	srv := &http.Server{
		Addr:         HTTPPort,
		Handler:      server.NewReverseProxy(container),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
	}
	srv.RegisterOnShutdown(cancel)

	return srv
}

func newGRPCServer(container *container.Container) *grpc.Server {
	s := grpc.NewServer()
	nodes.RegisterManagerServer(s, manager.NewManager(container))

	return s
}
