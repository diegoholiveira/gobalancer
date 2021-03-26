package container

import (
	"log"
	"time"

	"github.com/diegoholiveira/gobalancer/nodes"
)

const (
	MaxElapsedTime = 300 * time.Millisecond
)

type status struct {
	node *nodes.NodeStatus
	sync time.Time
}

func (s *status) IsAvailable() bool {
	n := s.node
	if n.ActiveConnections >= n.MaxConnections/2 {
		log.Printf("[MANAGER] Node too busy: %d\n", n.Port)

		return false
	}

	lastSync := time.Since(s.sync)

	if lastSync >= MaxElapsedTime {
		log.Printf("[MANAGER] Node unavailable after 300ms: %d\n", n.Port)

		return false
	}

	return true
}
