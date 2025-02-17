package client

import (
	"bytes"
	"encoding/json"
)

func (c *Client) GetAiAdapters() (*AiAdaptersResponse, error) {
	resp_body, err := c.httpRequest("/ai/adapters", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	response := &AiAdaptersResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}
