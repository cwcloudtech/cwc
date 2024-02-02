package admin

import (
	"bytes"
	"cwc/utils"
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
	if nil != err {
		return nil, err
	}
	resp_body, err := c.httpRequest(fmt.Sprintf("/admin/registry/%s/%s/provision", c.provider, c.region), "POST", buf)
	if nil != err {
		return nil, err
	}
	created_registry := &Registry{}
	err = json.NewDecoder(resp_body).Decode(created_registry)
	if nil != err {
		return nil, err
	}
	return created_registry, nil
}

func (c *Client) GetAllRegistries() (*[]Registry, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/registry/%s/%s/all", c.provider, c.region), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	registries := []Registry{}
	err = json.NewDecoder(body).Decode(&registries)

	if nil != err {
		return nil, err
	}
	return &registries, nil
}

func (c *Client) GetRegistry(registry_id string) (*Registry, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/registry/%s", registry_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	registry := &Registry{}
	err = json.NewDecoder(body).Decode(registry)
	if nil != err {
		return nil, err
	}
	return registry, nil
}

func (c *Client) UpdateRegistry(id string, args ...string) error {
	buf := bytes.Buffer{}
	var email string
	var updateCreds bool = true
	var renew RenewCredentials

	if len(args) > 0 {
		email = args[0]
		if !utils.IsValidEmail(email) {
			return fmt.Errorf("invalid email address")
		}
		renew = RenewCredentials{
			Email: email,
		}
	} else {
		renew = RenewCredentials{
			UpdateCreds: updateCreds,
		}
	}

	encode_err := json.NewEncoder(&buf).Encode(renew)
	if nil != encode_err {
		return encode_err
	}

	_, err := c.httpRequest(fmt.Sprintf("/admin/registry/%s", id), "PATCH", buf)
	if nil != err {
		return err
	}

	return nil
}

func (c *Client) DeleteRegistry(registry_id string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/registry/%s", registry_id), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}

	return nil
}
