package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) AdminAddEnvironment(name string, path string, roles []string, main_role string, is_private bool, description string, subdomains []string) (*Environment, error) {
	buf := bytes.Buffer{}
	environment := AddEnvironmentType{
		Name:        name,
		Path:        path,
		Roles:       roles,
		MainRole:    main_role,
		IsPrivate:   is_private,
		Description: description,
		SubDomains:  subdomains,
	}
	println(environment.Roles[1])
	err := json.NewEncoder(&buf).Encode(environment)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/admin/environment"), "POST", buf)
	if err != nil {
		return nil, err
	}
	created_env := &Environment{}
	err = json.NewDecoder(respBody).Decode(created_env)
	if err != nil {
		return nil, err
	}
	return created_env, nil
}

func (c *Client) AdminDeleteEnvironment(env_id string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/environment/%s", env_id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAllEnvironments() (*[]Environment, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/environment/all"), "GET", bytes.Buffer{})
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
	body, err := c.httpRequest(fmt.Sprintf("/admin/environment/%s", env_id), "GET", bytes.Buffer{})
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
