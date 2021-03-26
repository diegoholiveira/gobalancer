package manager

import (
	"context"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"github.com/diegoholiveira/gobalancer/nodes"
	"github.com/diegoholiveira/gobalancer/nodes/container"
)

type Manager struct {
	nodes.UnimplementedManagerServer

	container *container.Container
}

func NewManager(c *container.Container) *Manager {
	return &Manager{
		container: c,
	}
}

func (manager *Manager) SetStatus(ctx context.Context, node *nodes.NodeStatus) (*emptypb.Empty, error) {
	manager.container.SetStatus(node)

	return &emptypb.Empty{}, nil
}
