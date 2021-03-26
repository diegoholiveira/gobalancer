package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"

	"github.com/diegoholiveira/gobalancer/nodes"
)

const (
	balancer   string = ":8000"
	numServers int    = 10
	port       int    = 9000
)

type handler struct {
	mu      *sync.Mutex
	id      string
	ip      string
	port    uint32
	counter int32
}

func (h *handler) add(i int32) {
	h.mu.Lock()
	h.counter += i
	h.mu.Unlock()
}

func (h *handler) status() *nodes.NodeStatus {
	h.mu.Lock()
	defer h.mu.Unlock()

	return &nodes.NodeStatus{
		ID:                h.id,
		IP:                h.ip,
		Port:              h.port,
		MaxConnections:    100,
		ActiveConnections: uint32(h.counter),
	}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	n := rand.Intn(25)

	h.add(int32(n))
	defer h.add(int32(-n))

	time.Sleep(time.Duration(n) * time.Millisecond)

	fmt.Fprintf(w, "Pong from %s", h.id)
}

func newHTTPServer(ctx context.Context, cancel context.CancelFunc, port int) (*http.Server, *handler) {
	HTTPPort := fmt.Sprintf(":%d", port)

	nodeID, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("Impossible to generate an ID for this node")
	}

	h := &handler{
		mu:      &sync.Mutex{},
		id:      nodeID.String(),
		ip:      "127.0.0.1",
		port:    uint32(port),
		counter: 0,
	}

	srv := &http.Server{
		Handler:      h,
		Addr:         HTTPPort,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
	}
	srv.RegisterOnShutdown(cancel)

	return srv, h
}

func sendStatus(c nodes.ManagerClient, node *nodes.NodeStatus) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.SetStatus(ctx, node)
	if err != nil {
		log.Fatalf("could not send the node status: %v", err)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(
		shutdownSignal,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	conn, err := grpc.Dial(balancer, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	nodesManager := nodes.NewManagerClient(conn)

	log.Printf("Initialize %d http servers\n", numServers)

	servers := make([]*http.Server, numServers)
	handlers := make([]*handler, numServers)

	for i := 0; i < numServers; i++ {
		srv, handler := newHTTPServer(ctx, cancel, port+i)

		servers[i] = srv
		handlers[i] = handler

		go func() {
			log.Printf("http server at %d\n", handler.port)

			if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				log.Fatalf("http server error: %v", err)
			}
		}()

		sendStatus(nodesManager, handler.status())
	}

	ticker := time.NewTicker(10 * time.Millisecond)
	randomKiller := time.NewTicker(30 * time.Second)

	// number of interactions ignored
	j := 100
	// random node to ignore status
	n := -1

out:
	for {
		select {
		case <-shutdownSignal:
			ticker.Stop()

			break out
		case <-randomKiller.C:
			if 0 >= j {
				j = 100
				n = -1
			}

			if n == -1 {
				n = rand.Intn(numServers)

				handler := handlers[n]

				log.Printf("[KILLER] Node not sending status: %d\n", handler.port)
			}
		case <-ticker.C:
			for i := 0; i < numServers; i++ {
				if n == i {
					j-- // this counter will reset the random killer

					continue
				}

				handler := handlers[i]

				sendStatus(nodesManager, handler.status())
			}
		}
	}

	conn.Close()

	for i := 0; i < numServers; i++ {
		srv := servers[i]
		handler := handlers[i]

		gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()

		if err := srv.Shutdown(gracefullCtx); err != nil {
			log.Printf("http server at %d has a shutdown error: %v\n", handler.port, err)
		} else {
			log.Printf("http server at %d gracefully stopped\n", handler.port)
		}
	}
}
