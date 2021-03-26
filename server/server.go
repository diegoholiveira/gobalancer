package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const (
	DefaultIdleConnTimeout     time.Duration = 30 * time.Second
	DefaultMaxIdleConns        int           = 32
	DefaultMaxIdleConnsPerHost int           = 16
)

type (
	NodesContainer interface {
		HasNodes() bool
		Next() *url.URL
	}
)

func NewReverseProxy(container NodesContainer) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: NewDirector(container),
		Transport: NewRoundTripper(container, &http.Transport{
			MaxIdleConns:        DefaultMaxIdleConns,
			IdleConnTimeout:     DefaultIdleConnTimeout,
			MaxIdleConnsPerHost: DefaultMaxIdleConnsPerHost,
		}),
	}
}
