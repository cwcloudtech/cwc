package client

import (
	"bytes"
	"encoding/json"
)

func (c *Client) GetModels() (*ModelsResponse, error) {
	resp_body, err := c.httpRequest("/ai/models", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	response := &ModelsResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}
