package client

import (
	"bytes"
	"encoding/json"
)

func (c *Client) GetModels() (*ModelsResponse, error) {
	respBody, err := c.httpRequest("/ai/models", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	response := &ModelsResponse{}
	err = json.NewDecoder(respBody).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}
