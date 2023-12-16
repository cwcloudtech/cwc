package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllInvocations() (*[]Invocation, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/faas/invocations"), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := InvocationsResponse{}
	err = json.NewDecoder(body).Decode(&response)

	if nil != err {
		return nil, err
	}
	return &response.Results, nil
}

func (c *Client) GetInvocationInvokerById(id string) (*InvocationInvoker, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/faas/invocation/%s/invoker", id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := InvocationInvoker{}
	err = json.NewDecoder(body).Decode(&response)

	if nil != err {
		return nil, err
	}

	return &response, nil
}
