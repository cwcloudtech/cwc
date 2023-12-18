package client

import (
	"bytes"
	"cwc/config"
	"cwc/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewClient() (*Client, error) {
	region := config.GetDefaultRegion()
	provider := config.GetDefaultProvider()
	err := error(nil)

	if utils.IsBlank(provider) {
		err = fmt.Errorf("default provider is not set")
	} else if utils.IsBlank(region) {
		err = fmt.Errorf("default region is not set")
	}

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

	_, err = c.httpRequest(fmt.Sprintf("/api_keys/verify"), "POST", buf)
	if nil != err {
		return err
	}

	config.AddUserCredentials(access_key, secret_key)
	return nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if nil != err {
		return nil, err
	}

	user_token := config.GetUserToken()

	req.Header.Set("X-Auth-Token", user_token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if nil != err {
		return nil, err
	}

	switch {
	case resp.StatusCode >= 200 && resp.StatusCode < 400:
		return resp.Body, nil
	case resp.StatusCode >= 400 && resp.StatusCode < 500:
		resp_body := new(bytes.Buffer)
		_, err := resp_body.ReadFrom(resp.Body)
		if nil != err {
			return nil, fmt.Errorf("an error occurred")
		}

		errorResponse := ErrorResponse{}
		json.NewDecoder(resp_body).Decode(&errorResponse)
		if utils.IsBlank(errorResponse.Error) {
			return nil, fmt.Errorf("client error with status %d", resp.StatusCode)
		} else {
			return nil, fmt.Errorf(errorResponse.Error)
		}
	case resp.StatusCode >= 500:
		resp_body := new(bytes.Buffer)
		_, err := resp_body.ReadFrom(resp.Body)
		if nil != err {
			return nil, fmt.Errorf("an error occurred")
		}

		errorResponse := ErrorResponse{}
		json.NewDecoder(resp_body).Decode(&errorResponse)
		if utils.IsBlank(errorResponse.Error) {
			return nil, fmt.Errorf("server error with status %d", resp.StatusCode)
		} else {
			return nil, fmt.Errorf(errorResponse.Error)
		}
	}

	return nil, fmt.Errorf("unhandled status code: %d", resp.StatusCode)
}

func (c *Client) requestPath(path string) string {
	default_api_version := "v1"
	hostname := config.GetDefaultEndpoint()
	return fmt.Sprintf("%s/%s%s", hostname, default_api_version, path)
}
