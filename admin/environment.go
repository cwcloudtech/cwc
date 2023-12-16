package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) AdminAddEnvironment(environment Environment) (*Environment, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(environment)
	if nil != err {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/admin/environment"), "POST", buf)
	if nil != err {
		return nil, err
	}
	created_env := &Environment{}
	err = json.NewDecoder(respBody).Decode(created_env)
	if nil != err {
		return nil, err
	}
	return created_env, nil
}

func (c *Client) AdminDeleteEnvironment(env_id string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/environment/%s", env_id), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}

func (c *Client) GetAllEnvironments() (*[]Environment, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/environment/all"), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	environments := []Environment{}
	err = json.NewDecoder(body).Decode(&environments)

	if nil != err {
		return nil, err
	}
	return &environments, nil
}

func (c *Client) GetEnvironment(env_id string) (*Environment, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/environment/%s", env_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	environment := &Environment{}
	err = json.NewDecoder(body).Decode(environment)
	if nil != err {
		return nil, err
	}
	return environment, nil
}
