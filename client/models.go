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

func (c *Client) LoadModel(model string) (*LoadModelResponse, error) {
	resp_body, err := c.httpRequest("/ai/model/"+model, "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	response := &LoadModelResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}
