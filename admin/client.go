package admin

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

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		resp_body := new(bytes.Buffer)
		_, err := resp_body.ReadFrom(resp.Body)
		if nil != err {
			return nil, fmt.Errorf("an error occured")
		}

		errorResponse := ErrorResponse{}
		json.NewDecoder(resp_body).Decode(&errorResponse)
		if utils.IsBlank(errorResponse.Error) {
			return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
		} else {
			return nil, fmt.Errorf(errorResponse.Error)
		}
	}

	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	default_api_version := "v1"
	hostname := config.GetDefaultEndpoint()
	return fmt.Sprintf("%s/%s%s", hostname, default_api_version, path)
}
