package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllRegistries() (*[]Registry, error) {
	body, err := c.httpRequest(fmt.Sprintf("/registry/%s/%s", c.provider, c.region), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	registries := []Registry{}
	err = json.NewDecoder(body).Decode(&registries)

	if err != nil {
		return nil, err
	}
	return &registries, nil
}

func (c *Client) GetRegistry(registry_id string) (*Registry, error) {
	body, err := c.httpRequest(fmt.Sprintf("/registry/%s/%s/%s", c.provider, c.region, registry_id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	registry := &Registry{}
	err = json.NewDecoder(body).Decode(registry)
	if err != nil {
		return nil, err
	}
	return registry, nil
}

func (c *Client) UpdateRegistry(id string) error {
	buf := bytes.Buffer{}

	_, err := c.httpRequest(fmt.Sprintf("/registry/%s/%s/%s", c.provider, c.region, id), "PATCH", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteRegistry(registry_id string) error {
	_, err := c.httpRequest(fmt.Sprintf("/registry/%s/%s/%s", c.provider, c.region, registry_id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
