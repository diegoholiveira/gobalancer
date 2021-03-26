package server

import "net/http"

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func NewRoundTripper(container NodesContainer, original http.RoundTripper) http.RoundTripper {
	if original == nil {
		original = http.DefaultTransport
	}

	return roundTripperFunc(func(request *http.Request) (*http.Response, error) {
		if !container.HasNodes() {
			return &http.Response{
				Proto:      "HTTP/1.1",
				ProtoMajor: 1,
				ProtoMinor: 1,
				StatusCode: http.StatusServiceUnavailable,
				Body:       http.NoBody,
			}, nil
		}
		return original.RoundTrip(request)
	})
}
