package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllFunctions() (*[]Function, error) {
	body, err := c.httpRequest("/admin/faas/functions", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := FunctionsResponse{}
	err = json.NewDecoder(body).Decode(&response)

	if nil != err {
		return nil, err
	}

	return &response.Results, nil
}

func (c *Client) GetFunctionOwnerById(id string) (*FunctionOwner, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/faas/function/%s/owner", id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	response := FunctionOwner{}
	err = json.NewDecoder(body).Decode(&response)

	if nil != err {
		return nil, err
	}

	return &response, nil
}
