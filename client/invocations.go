package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllInvocations() (*[]Invocation, error) {
	body, err := c.httpRequest("/faas/invocations", "GET", bytes.Buffer{})
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

func (c *Client) GetInvocationById(invocation_id string) (*Invocation, error) {
	body, err := c.httpRequest(fmt.Sprintf("/faas/invocation/%s", invocation_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	invocation := &Invocation{}
	err = json.NewDecoder(body).Decode(invocation)
	if nil != err {
		return nil, err
	}
	return invocation, nil
}

func (c *Client) AddInvocation(content InvocationAddContent, synchronous bool) (*Invocation, error) {
	buf := bytes.Buffer{}
	if len(content.Args) == 0 {
		content.Args = []Argument{}
	}

	invocation := &AddInvocationBody{
		Content: content,
	}

	err := json.NewEncoder(&buf).Encode(invocation)
	if nil != err {
		return nil, err
	}

	endpoint := "/faas/invocation"
	if synchronous {
		endpoint = endpoint + "/sync"
	}

	respBody, err := c.httpRequest(endpoint, "POST", buf)
	if nil != err {
		return nil, err
	}

	created_invocation := &Invocation{}
	err = json.NewDecoder(respBody).Decode(created_invocation)
	if nil != err {
		return nil, err
	}

	return created_invocation, nil
}

func (c *Client) DeleteInvocationById(invocationId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/faas/invocation/%s", invocationId), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}

	return nil
}

func (c *Client) TruncateInvocations() error {
	_, err := c.httpRequest("/faas/invocations", "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}

	return nil
}
