package http

import (
	"net/http"
)

//go:generate mockgen -package=automock -source=client.go -destination=./tests/gomock/automock/http_client_mock.go HTTPClient

type (
	// ClientWithBasicAuth implements client with basic authorization
	ClientWithBasicAuth struct {
		httpClient HTTPClient
		user, pass string
	}

	// HTTPClient is an interface for testing a request object.
	HTTPClient interface {
		Do(*http.Request) (*http.Response, error)
	}
)

// NewClientWithBasicAuth returns httpClient which always enriches request with basic auth
func NewClientWithBasicAuth(user, pass string) *ClientWithBasicAuth {
	return &ClientWithBasicAuth{
		user: user,
		pass: pass,
		httpClient: http.DefaultClient,
	}
}

// Do executes passed HTTP request and returns a response. It enriches request with basic auth.
func (c *ClientWithBasicAuth) Do(req *http.Request) (*http.Response, error) {
	req.SetBasicAuth(c.user, c.pass)
	return c.httpClient.Do(req)
}

func (c *ClientWithBasicAuth) WithClient(client HTTPClient) *ClientWithBasicAuth {
	c.httpClient = client
	return c
}
