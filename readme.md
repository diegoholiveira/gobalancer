# gobalancer

gobalancer is an experimental load balancer.

In this repository we have three binaries:

- cmd/server : the load balancer
- cmd/node   : instance a predefined number of nodes
- cmd/test   : send a predefined number of requests into the load balancer

The nodes must to send they current status to the load balancer to say, hey, I'm alive.
The server will recicle the nodes to remove those who're too busy or don't send they status frequently.

The use gRPC to send the node status to server. And expose the http rest interface in other server.
