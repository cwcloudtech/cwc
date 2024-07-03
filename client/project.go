package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) AddProject(project_name string, host string, token string, git_username string, namespace string, project_type string) (*Project, error) {
	buf := bytes.Buffer{}
	project := AddProjectBody{
		Name:        project_name,
		Host:        host,
		Token:       token,
		GitUsername: git_username,
		Namespace:   namespace,
		Type:        project_type,
	}

	err := json.NewEncoder(&buf).Encode(project)
	if nil != err {
		return nil, err
	}
	resp_body, err := c.httpRequest("/project", "POST", buf)
	if nil != err {
		return nil, err
	}
	created_project := &Project{}
	err = json.NewDecoder(resp_body).Decode(created_project)
	if nil != err {
		return nil, err
	}
	return created_project, nil
}

func (c *Client) DeleteProjectById(projectId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/project/%s", projectId), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}

func (c *Client) DeleteProjectByName(projectName string) error {
	_, err := c.httpRequest(fmt.Sprintf("/project/name/%s", projectName), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}
func (c *Client) DeleteProjectByUrl(projectUrl string) error {
	encodedUrl := url.QueryEscape(projectUrl)
	_, err := c.httpRequest(fmt.Sprintf("/project/url/%s", encodedUrl), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}
func (c *Client) GetAllProjects(projectType string) (*[]Project, error) {
	if projectType == "" {
		projectType = "vm"
	}
	body, err := c.httpRequest(fmt.Sprintf("/project?type=%s", projectType), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	projects := []Project{}
	err = json.NewDecoder(body).Decode(&projects)

	if nil != err {
		return nil, err
	}
	return &projects, nil
}

func (c *Client) GetProjectById(project_id string) (*Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/project/%s", project_id), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	project := &Project{}
	err = json.NewDecoder(body).Decode(project)
	if nil != err {
		return nil, err
	}
	return project, nil
}

func (c *Client) GetProjectByName(project_name string) (*Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/project/name/%s", project_name), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	project := &Project{}
	err = json.NewDecoder(body).Decode(project)
	if nil != err {
		return nil, err
	}
	return project, nil
}

func (c *Client) GetProjectByUrl(project_url string) (*Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/project/url/%s", project_url), "GET", bytes.Buffer{})
	if nil != err {
		return nil, err
	}
	project := &Project{}
	err = json.NewDecoder(body).Decode(project)
	if nil != err {
		return nil, err
	}
	return project, nil
}
