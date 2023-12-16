package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) AdminAddProject(user_email string, project_name string, host string, token string, git_username string, namespace string) (*Project, error) {
	buf := bytes.Buffer{}
	project := AddProjectBody{
		Name:        project_name,
		Host:        host,
		Token:       token,
		Email:       user_email,
		GitUsername: git_username,
		Namespace:   namespace,
	}

	err := json.NewEncoder(&buf).Encode(project)
	if nil != err {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/admin/project"), "POST", buf)
	if nil != err {
		return nil, err
	}
	created_project := &Project{}
	err = json.NewDecoder(respBody).Decode(created_project)
	if nil != err {
		return nil, err
	}
	return created_project, nil
}

func (c *Client) AdminDeleteProjectById(projectId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/project/%s", projectId), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}

func (c *Client) AdminDeleteProjectByName(projectName string) error {
	_, err := c.httpRequest(fmt.Sprintf("/admin/project/name/%s", projectName), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}
func (c *Client) AdminDeleteProjectByUrl(projectUrl string) error {
	encodedUrl := url.QueryEscape(projectUrl)
	_, err := c.httpRequest(fmt.Sprintf("/admin/project/url/%s", encodedUrl), "DELETE", bytes.Buffer{})
	if nil != err {
		return err
	}
	return nil
}
func (c *Client) AdminGetAllProjects() (*[]Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/project"), "GET", bytes.Buffer{})
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

func (c *Client) AdminGetProjectById(project_id string) (*Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/project/%s", project_id), "GET", bytes.Buffer{})
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

func (c *Client) AdminGetProjectByName(project_name string) (*Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/project/name/%s", project_name), "GET", bytes.Buffer{})
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

func (c *Client) AdminGetProjectByUrl(project_url string) (*Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/admin/project/url/%s", project_url), "GET", bytes.Buffer{})
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
