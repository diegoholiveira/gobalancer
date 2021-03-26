package container

import (
	"fmt"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/diegoholiveira/gobalancer/nodes"
)

type Container struct {
	// status keeps all status received
	status *sync.Map
	// resolver calculates the next node to be used
	resolver *urlResolver
	// registered keeps the number of nodes registered
	registered uint32

	mu *sync.Mutex
}

func NewContainer() *Container {
	return &Container{
		status:   &sync.Map{},
		resolver: newURLResolver(make([]*url.URL, 0)),
		mu:       &sync.Mutex{},
	}
}

func (c *Container) SetStatus(n *nodes.NodeStatus) {
	url, _ := url.Parse(fmt.Sprintf("http://%s:%d", n.IP, n.Port))
	c.status.Store(n.ID, &status{
		node: n,
		sync: time.Now(),
	})
	c.resolver.add(url)
	_ = atomic.AddUint32(&c.registered, 1)
}

func (c *Container) HasNodes() bool {
	return c.registered > 0
}

func (c *Container) Next() *url.URL {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.resolver.next()
}

// Update removes the nodes unavailable or too busy
func (c *Container) Update() {
	available := make([]*url.URL, 0)
	removed := 0
	registered := 0
	c.status.Range(func(_id interface{}, value interface{}) bool {
		var s = value.(*status)
		if !s.IsAvailable() {
			c.status.Delete(_id)
			removed += 1
			return true
		}

		url, _ := url.Parse(fmt.Sprintf("http://%s:%d", s.node.IP, s.node.Port))
		available = append(available, url)
		registered += 1

		return true

	})

	c.mu.Lock()
	c.registered = uint32(registered)
	c.resolver = newURLResolver(available)
	c.mu.Unlock()
}
