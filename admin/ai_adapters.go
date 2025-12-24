package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllAIAdapters() (*AdminAIAdaptersResponse, error) {
	resp_body, err := c.httpRequest("/admin/ai/adapters", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	response := &AdminAIAdaptersResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}

func (c *Client) GetAIAdapterById(adapterId string) (*AdminAIAdapterDetailResponse, error) {
	resp_body, err := c.httpRequest(fmt.Sprintf("/admin/ai/adapters/%s", adapterId), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	response := &AdminAIAdapterDetailResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}

func (c *Client) CreateAIAdapter(adapter AdminAIAdapterRequest) (*AdminAIAdapterResponse, error) {
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

	resp_body, err := c.httpRequest("/admin/ai/adapters", "POST", buf)
	if nil != err {
		return nil, err
	}

	response := &AdminAIAdapterResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}

func (c *Client) UpdateAIAdapter(adapterId string, adapter AdminAIAdapterRequest) (*AdminAIAdapterResponse, error) {
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

	resp_body, err := c.httpRequest(fmt.Sprintf("/admin/ai/adapters/%s", adapterId), "PUT", buf)
	if nil != err {
		return nil, err
	}

	response := &AdminAIAdapterResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}

func (c *Client) DeleteAIAdapter(adapterId string) (*AdminAIAdapterResponse, error) {
	resp_body, err := c.httpRequest(fmt.Sprintf("/admin/ai/adapters/%s", adapterId), "DELETE", bytes.Buffer{})
	if nil != err {
		return nil, err
	}

	response := &AdminAIAdapterResponse{}
	err = json.NewDecoder(resp_body).Decode(response)
	if nil != err {
		return nil, err
	}

	return response, nil
}
