package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) AdminGetAllInstances() (*[]Instance, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/instance/%s/%s/all", c.provider, c.region), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	instances := []Instance{}
	err = json.NewDecoder(body).Decode(&instances)

	if err != nil {
		return nil, err
	}
	return &instances, nil
}

func (c *Client) GetInstance(instance_id string) (*Instance, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/instance/%s", instance_id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	instance := &Instance{}
	err = json.NewDecoder(body).Decode(instance)
	if err != nil {
		return nil, err
	}
	return instance, nil
}


func (c *Client) AdminAddInstance(user_email string, instance_name string, project_id int, project_name string, project_url string, instance_size string, environment string, zone string, dns_zone string) (*Instance, error) {
	buf := bytes.Buffer{}
	instance := Instance{
		Name:          instance_name,
		Zone:          zone,
		Instance_type: instance_size,
		Root_dns_zone: dns_zone,
		Environment:   environment,
		Project:       project_id,
		Email:         user_email,
		Project_name:  project_name,
		Project_url:   project_url,
		Region:        c.region,
	}

	err := json.NewEncoder(&buf).Encode(instance)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/admin/instance/%s/%s/%s/provision/%s", c.provider, c.region, instance.Zone, instance.Environment), "POST", buf)
	if err != nil {
		return nil, err
	}
	created_instance := &Instance{}
	err = json.NewDecoder(respBody).Decode(created_instance)
	if err != nil {
		return nil, err
	}
	return created_instance, nil
}

func (c *Client) AdminUpdateInstance(id string, status string) error {
	buf := bytes.Buffer{}

	UpdateInstanceRequest := &UpdateInstanceRequest{
		Status: status,
	}
	err := json.NewEncoder(&buf).Encode(UpdateInstanceRequest)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("/admin/instance/%s", id), "PATCH", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AdminDeleteInstance(instanceId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/instance/%s", instanceId), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AdminRefreshInstance(instanceId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/instance/%s/refresh", instanceId), "POST", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}
