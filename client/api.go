package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}


type Client struct {
	region   string
	authToken  string
	httpClient *http.Client
}
type Project struct {
	Id int `json:"id"`
	Name string `json:"name"` 
	Gitlab_url string `json:"gitlab_project_url"`
	Instance_type string `json:"type"`
	Environment string `json:"environment"`
	Status string `json:"status"`
	Email string `json:"email"`
	Region string `json:"region"`
	Ip_address string `json:"ip_address"`

}

func NewClient(region string, token string) *Client {
	return &Client{
		region:       region,
		authToken:  token,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetAll() (*[]Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/instance/%s",c.region), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	projects := []Project{}
	err = json.NewDecoder(body).Decode(&projects)
	
	if err != nil {
		print(err.Error())
		return nil, err
	}
	return &projects, nil
}

func (c *Client) GetProject(project_id string) (*Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s",c.region, project_id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	project := &Project{}
	err = json.NewDecoder(body).Decode(project)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (c *Client) AddProject(project_name string, instance_size string,environment string, email string) (*Project,error) {
	buf := bytes.Buffer{}
	project := Project{
		Name:project_name,
		Instance_type:instance_size,
		Environment:environment,
		Email: email,
		Region:c.region,
	}

	err := json.NewEncoder(&buf).Encode(project)
	if err != nil {
		fmt.Print("err1")
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/instance/%s/provision/%s",c.region, project.Environment), "POST", buf)
	if err != nil {
		fmt.Print("err2")
		return nil, err
	}
	created_project := &Project{}
	err = json.NewDecoder(respBody).Decode(created_project)
	if err != nil {
		fmt.Print("err3")
		return nil, err
	}
	return created_project, nil
}

func (c *Client) UpdateProject(project *Project) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(project)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("/instance/%s/%s", c.region,project.Environment), "PATCH", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteProject(itemName string) error {
	_, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s",c.region, itemName), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)

	req.Header.Set("X-User-Token", c.authToken)


	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		fmt.Println(respBody.String())
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			fmt.Println("err!!")
			return nil, fmt.Errorf("an error occured")
		}
		fmt.Println("err!")
		errorResponse := ErrorResponse{}
		json.NewDecoder(respBody).Decode(&errorResponse)
		fmt.Println("------")
		fmt.Println(errorResponse.Error)
		return nil, fmt.Errorf(errorResponse.Error)
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	hostname := "https://cloud-api.comwork.io/v1"
	return fmt.Sprintf("%s%s", hostname, path)
}
