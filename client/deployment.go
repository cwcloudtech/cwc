package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (c *Client) GetAllDeployments() (*[]Deployment, error) {
	body, err := c.httpRequest("/kubernetes/deployment", "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	var response []Deployment
	err = json.NewDecoder(body).Decode(&response)

	if nil != err {
		return nil, err
	}

	return &response, nil
}

func (c *Client) GetDeploymentById(deployment_id string) (*DeploymentByIdResponse, error) {
	body, err := c.httpRequest(fmt.Sprintf("/kubernetes/deployment/%s", deployment_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	deployment := &DeploymentByIdResponse{}
	err = json.NewDecoder(body).Decode(deployment)
	if nil != err {
		return nil, err
	}
	return deployment, nil
}

func (c *Client) DeleteDeploymentById(deploymentId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/kubernetes/deployment/%s", deploymentId), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}

func (c *Client) CreateDeployment(deployment CreationDeployment) (*CreationDeployment, error) {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(deployment)
	if nil != err {
		return nil, err
	}
	resp_body, err := c.httpRequest("/kubernetes/deployment", "POST", buf)
	if nil != err {
		return nil, err
	}
	created_deployment := &CreationDeployment{}
	err = json.NewDecoder(resp_body).Decode(created_deployment)
	if nil != err {
		return nil, err
	}
	return created_deployment, nil
}
