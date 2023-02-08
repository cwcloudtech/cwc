package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) AdminAddRegistry(user_email string, name string, reg_type string) (*Registry, error) {
	buf := bytes.Buffer{}
	registry := Registry{
		Name:  name,
		Type:  reg_type,
		Email: user_email,
	}

	err := json.NewEncoder(&buf).Encode(registry)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/admin/registry/%s/%s/provision", c.provider, c.region), "POST", buf)
	if err != nil {
		return nil, err
	}
	created_registry := &Registry{}
	err = json.NewDecoder(respBody).Decode(created_registry)
	if err != nil {
		return nil, err
	}
	return created_registry, nil
}

func (c *Client) GetAllRegistries() (*[]Registry, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/registry/%s/%s/all", c.provider, c.region), "GET", bytes.Buffer{})
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
	body, err := c.httpRequest(fmt.Sprintf("/admin/registry/%s", registry_id), "GET", bytes.Buffer{})
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

func (c *Client) UpdateRegistry(id string, email string) error {
	buf := bytes.Buffer{}
	var updateCreds bool = true

	renew := RenewCredentials{
		Email: email,
		UpdateCreds: updateCreds,
	}

	if len(email) > 0 && valid(email) {
		updateCreds = false
	}

	encode_err := json.NewEncoder(&buf).Encode(renew)
	if encode_err != nil {
		return encode_err
	}

	_, err := c.httpRequest(fmt.Sprintf("/admin/registry/%s", id), "PATCH", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteRegistry(registry_id string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/registry/%s", registry_id), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
