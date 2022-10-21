package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllEnvironments() (*[]Environment, error) {
	body, err := c.httpRequest(fmt.Sprintf("/environment/all"), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	environments := []Environment{}
	err = json.NewDecoder(body).Decode(&environments)

	if err != nil {
		return nil, err
	}
	return &environments, nil
}
