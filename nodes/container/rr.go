package container

import (
	"math/rand"
	"net/url"
	"sync"
)

type urlResolver struct {
	mu       *sync.Mutex
	nextNode uint32

	// available have all hosts available
	available []*url.URL
}

func newURLResolver(available []*url.URL) *urlResolver {
	return &urlResolver{
		mu:        &sync.Mutex{},
		available: available,
	}
}

func (rr *urlResolver) add(url *url.URL) {
	rr.mu.Lock()
	rr.available = append(rr.available, url)
	rr.mu.Unlock()
}

func (rr *urlResolver) next() *url.URL {
	rr.mu.Lock()
	n := rand.Intn(len(rr.available))
	target := rr.available[n]
	rr.mu.Unlock()
	return &url.URL{
		Scheme:      target.Scheme,
		Opaque:      target.Opaque,
		User:        target.User,
		Host:        target.Host,
		Path:        target.Path,
		RawPath:     target.RawPath,
		ForceQuery:  target.ForceQuery,
		RawQuery:    target.RawQuery,
		Fragment:    target.Fragment,
		RawFragment: target.RawFragment,
	}
}
