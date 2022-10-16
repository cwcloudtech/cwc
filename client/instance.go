package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllInstances() (*[]Instance, error) {
	body, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s", c.provider, c.region), "GET", bytes.Buffer{})
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
	body, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s/%s", c.provider, c.region, instance_id), "GET", bytes.Buffer{})
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

func (c *Client) AddInstance(instance_name string, project_id int, instance_size string, environment string, zone string, dns_zone string) (*Instance, error) {
	buf := bytes.Buffer{}
	instance := Instance{
		Name:          instance_name,
		Zone:          zone,
		Instance_type: instance_size,
		Root_dns_zone: dns_zone,
		Environment:   environment,
		Project:       project_id,
		Region:        c.region,
	}

	err := json.NewEncoder(&buf).Encode(instance)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s/%s/provision/%s", c.provider, c.region, instance.Zone, instance.Environment), "POST", buf)
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

func (c *Client) UpdateInstance(id string, status string) error {
	buf := bytes.Buffer{}

	UpdateInstanceRequest := &UpdateInstanceRequest{
		Status: status,
	}
	err := json.NewEncoder(&buf).Encode(UpdateInstanceRequest)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("/instance/%s/%s/%s", c.provider, c.region, id), "PATCH", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteInstance(instanceId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s/%s", c.provider, c.region, instanceId), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AttachInstance(project_id int, playbook string, instance_size string) (*Instance, error) {
	buf := bytes.Buffer{}
	instance := AttachInstanceRequest{
		Name:          playbook,
		ProjectId:     project_id,
		Instance_type: instance_size,
	}

	err := json.NewEncoder(&buf).Encode(instance)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/instance/%s/attach/%v", c.region, instance.ProjectId), "POST", buf)
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
