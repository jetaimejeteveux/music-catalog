package httpclient

import "net/http"

//go:generate mockgen -source=client.go -destination=client_mock.go -package=httpclient
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient HttpClient
}

func NewClient(httpClient HttpClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}
