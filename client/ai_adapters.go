package client

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (c *Client) GetExternalAIAdapters() (*[]AIAdapter, error) {
	resp_body, err := c.httpRequest("/ai/adapters/external", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	var adapters []AIAdapter
	err = json.NewDecoder(resp_body).Decode(&adapters)
	if nil != err {
		return nil, err
	}

	return &adapters, nil
}

func (c *Client) GetAIAdapterById(adapterId string) (*AIAdapter, error) {
	resp_body, err := c.httpRequest(fmt.Sprintf("/ai/adapters/%s", adapterId), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	adapter := &AIAdapter{}
	err = json.NewDecoder(resp_body).Decode(adapter)
	if nil != err {
		return nil, err
	}

	return adapter, nil
}

func (c *Client) CreateAIAdapter(adapter AIAdapterRequest) (*AIAdapterResponse, error) {
	if adapter.Timeout == 0 {
		adapter.Timeout = 30
	}
	if adapter.Headers == nil {
		adapter.Headers = []AIAdapterHeader{}
	}

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(adapter)
	if nil != err {
		return nil, err
	}

	resp_body, err := c.httpRequest("/ai/adapters", "POST", buf)
	if nil != err {
		return nil, err
	}

	response := &AIAdapterResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateAIAdapter(adapterId string, adapter AIAdapterRequest) (*AIAdapterResponse, error) {
	if adapter.Timeout == 0 {
		adapter.Timeout = 30
	}
	if adapter.Headers == nil {
		adapter.Headers = []AIAdapterHeader{}
	}

	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(adapter)
	if nil != err {
		return nil, err
	}

	resp_body, err := c.httpRequest(fmt.Sprintf("/ai/adapters/%s", adapterId), "PUT", buf)
	if nil != err {
		return nil, err
	}

	response := &AIAdapterResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteAIAdapter(adapterId string) (*AIAdapterResponse, error) {
	resp_body, err := c.httpRequest(fmt.Sprintf("/ai/adapters/%s", adapterId), "DELETE", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	response := &AIAdapterResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}
