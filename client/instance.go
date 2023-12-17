package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllInstances() (*[]Instance, error) {
	body, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s", c.provider, c.region), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	instances := []Instance{}
	err = json.NewDecoder(body).Decode(&instances)

	if nil != err {
		return nil, err
	}
	return &instances, nil
}

func (c *Client) GetInstance(instance_id string) (*Instance, error) {
	body, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s/%s", c.provider, c.region, instance_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	instance := &Instance{}
	err = json.NewDecoder(body).Decode(instance)
	if nil != err {
		return nil, err
	}
	return instance, nil
}

func (c *Client) AddInstance(instance_name string, project_id int, project_name string, project_url string, instance_size string, environment string, zone string, dns_zone string) (*Instance, error) {
	buf := bytes.Buffer{}
	instance := Instance{
		Name:          instance_name,
		Zone:          zone,
		Instance_type: instance_size,
		Root_dns_zone: dns_zone,
		Environment:   environment,
		Project:       project_id,
		Project_name:  project_name,
		Project_url:   project_url,
		Region:        c.region,
	}

	err := json.NewEncoder(&buf).Encode(instance)
	if nil != err {
		return nil, err
	}

	resp_body, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s/%s/provision/%s", c.provider, c.region, instance.Zone, instance.Environment), "POST", buf)
	if nil != err {
		return nil, err
	}

	created_instance := &Instance{}
	err = json.NewDecoder(resp_body).Decode(created_instance)
	if nil != err {
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
	if nil != err {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("/instance/%s/%s/%s", c.provider, c.region, id), "PATCH", buf)
	if nil != err {
		return err
	}
	return nil
}

func (c *Client) DeleteInstance(instanceId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s/%s", c.provider, c.region, instanceId), "DELETE", bytes.Buffer{})
	if nil != err {
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
	if nil != err {
		return nil, err
	}

	resp_body, err := c.httpRequest(fmt.Sprintf("/instance/%s/attach/%v", c.region, instance.ProjectId), "POST", buf)
	if nil != err {
		return nil, err
	}

	created_instance := &Instance{}
	err = json.NewDecoder(resp_body).Decode(created_instance)
	if nil != err {
		return nil, err
	}

	return created_instance, nil
}

func GetInstancesTypes() (*InstancesTypes, error) {
	c, err := NewClient()
	if nil != err {
		return nil, err
	}
	body, err := c.httpRequest(fmt.Sprintf("/%s/instance_types", c.provider), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	instancesTypes := &InstancesTypes{}
	err = json.NewDecoder(body).Decode(instancesTypes)
	if nil != err {
		fmt.Println(err.Error())
		return nil, err
	}
	return instancesTypes, nil
}
