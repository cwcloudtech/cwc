package admin

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

func (c *Client) GetEnvironment(env_id string) (*Environment, error) {
	body, err := c.httpRequest(fmt.Sprintf("/environment/%s", env_id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	environment := &Environment{}
	err = json.NewDecoder(body).Decode(environment)
	if err != nil {
		return nil, err
	}
	return environment, nil
}
