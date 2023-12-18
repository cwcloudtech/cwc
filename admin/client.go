package admin

import (
	"bytes"
	"cwc/httpcli"
	"io"
	"net/http"
)

func NewClient() (*Client, error) {
	region, provider, err := httpcli.ConfigClient()

	return &Client{
		region:     region,
		provider:   provider,
		httpClient: &http.Client{},
	}, err
}

func (c *Client) httpRequest(path string, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	return httpcli.HttpRequest(c.httpClient, path, method, body)
}
