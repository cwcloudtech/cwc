package client

import (
	"bytes"
	"cwc/config"
	"cwc/httpcli"
	"encoding/json"
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

func (c *Client) UserLogin(access_key string, secret_key string) error {
	buf := bytes.Buffer{}
	project := ApiKey{
		Accesskey: access_key,
		SecretKey: secret_key,
	}

	err := json.NewEncoder(&buf).Encode(project)
	if nil != err {
		return err
	}

	_, err = c.httpRequest("/api_keys/verify", "POST", buf)
	if nil != err {
		return err
	}

	config.AddUserCredentials(access_key, secret_key)
	return nil
}

func (c *Client) httpRequest(path string, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	return httpcli.HttpRequest(c.httpClient, path, method, body)
}
