package client

import (
	"net/http"
)

var (
	DefaultAPIURL = "https://pokeapi.co"
)

type Client struct {
	apiURL     string
	httpClient *http.Client
}

type Options func(c *Client)

func NewClient(options ...Options) *Client {
	client := &Client{
		apiURL:     DefaultAPIURL,
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		option(client)
	}
	return client
}

func withAPIURL(url string) Options {
	return func(c *Client) {
		c.apiURL = url
	}
}

func withHTTPClient(httpClient *http.Client) Options {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
