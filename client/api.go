package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type UpdateInstanceRequest struct {
	Status string `json:"status"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type ReponseLogin struct {
	Token     string `json:"token"`
	Confirmed bool   `json:"confirmed"`
}
type Client struct {
	region     string
	httpClient *http.Client
}

type Environment struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

type Instance struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Instance_type string `json:"type"`
	Environment   string `json:"environment"`
	Status        string `json:"status"`
	Project       int    `json:"project_id"`
	Region        string `json:"region"`
	Ip_address    string `json:"ip_address"`
}

type Project struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Url       string     `json:"url"`
	CreatedAt string     `json:"created_at"`
	Region    string     `json:"region"`
	Instances []Instance `json:"instances"`
}

func NewClient() *Client {
	region := GetDefaultRegion()
	return &Client{
		region:     region,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetAllInstances() (*[]Instance, error) {
	body, err := c.httpRequest(fmt.Sprintf("/instance/%s", c.region), "GET", bytes.Buffer{})
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
	body, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s", c.region, instance_id), "GET", bytes.Buffer{})
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

func (c *Client) AddInstance(instance_name string, project_id int, instance_size string, environment string) (*Instance, error) {
	buf := bytes.Buffer{}
	instance := Instance{
		Name:          instance_name,
		Instance_type: instance_size,
		Environment:   environment,
		Project:       project_id,
		Region:        c.region,
	}

	err := json.NewEncoder(&buf).Encode(instance)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/instance/%s/provision/%s", c.region, instance.Environment), "POST", buf)
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
	_, err = c.httpRequest(fmt.Sprintf("/instance/%s/%s", c.region, id), "PATCH", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteInstance(instanceId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s", c.region, instanceId), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddProject(project_name string) (*Project, error) {
	buf := bytes.Buffer{}
	project := Project{
		Name:   project_name,
		Region: c.region,
	}

	err := json.NewEncoder(&buf).Encode(project)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/project/%s", c.region), "POST", buf)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	created_project := &Project{}
	err = json.NewDecoder(respBody).Decode(created_project)
	if err != nil {
		fmt.Printf("hahah")
		return nil, err
	}
	return created_project, nil
}

func (c *Client) DeleteProject(projectId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/project/%s/%s", c.region, projectId), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAllProjects() (*[]Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/project/%s", c.region), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	projects := []Project{}
	err = json.NewDecoder(body).Decode(&projects)

	if err != nil {
		return nil, err
	}
	return &projects, nil
}

func (c *Client) GetAllEnvironments() (*[]Environment, error) {
	body, err := c.httpRequest(fmt.Sprintf("/environment"), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	environments := []Environment{}
	err = json.NewDecoder(body).Decode(&environments)

	if err != nil {
		return nil, err
	}
	return &environments, nil
}

func (c *Client) GetEnvironment(env_id string) (*Environment, error) {
	body, err := c.httpRequest(fmt.Sprintf("/environment/%s", env_id), "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}
	environment := &Environment{}
	err = json.NewDecoder(body).Decode(environment)
	if err != nil {
		return nil, err
	}
	return environment, nil
}

func (c *Client) GetProject(project_id string) (*Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/project/%s/%s", c.region, project_id), "GET", bytes.Buffer{})
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
func (c *Client) UserLogin(email string, password string) error {
	buf := bytes.Buffer{}

	loginBody := &LoginBody{
		Email:    email,
		Password: password,
	}
	err := json.NewEncoder(&buf).Encode(loginBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.requestPath("/auth/login"), &buf)
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return fmt.Errorf("an error occured")
		}
		errorResponse := ErrorResponse{}
		json.NewDecoder(respBody).Decode(&errorResponse)
		return fmt.Errorf(errorResponse.Error)
	}

	if err != nil {
		return err
	}
	reponseLogin := &ReponseLogin{}
	err = json.NewDecoder(resp.Body).Decode(reponseLogin)
	addUserToken(reponseLogin.Token)
	if err != nil {
		return err
	}
	return nil
}
func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)

	user_token, err := getUserToken()
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-User-Token", user_token)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("an error occured")
		}
		errorResponse := ErrorResponse{}
		json.NewDecoder(respBody).Decode(&errorResponse)
		return nil, fmt.Errorf(errorResponse.Error)
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	hostname := "https://cloud-api.comwork.io/v1"
	return fmt.Sprintf("%s%s", hostname, path)
}

func addUserToken(token string) {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)

	}
	if _, err := os.Stat(dirname + "/.cwc"); os.IsNotExist(err) {
		err := os.Mkdir(dirname+"/.cwc", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.Create(dirname + "/.cwc/credentials")
	if err != nil {
		log.Fatal(err)

	}
	_, err = f.WriteString("access_token = " + token)
	if err != nil {
		log.Fatal(err)

	}
}

func getUserToken() (string, error) {
	dirname, err := os.UserHomeDir()

	if err != nil {
		_err := errors.New("cwc: access denied, please login")
		return "", _err
	}

	content, err := ioutil.ReadFile(dirname + "/.cwc/credentials")
	if err != nil {
		_err := errors.New("cwc: access denied, please login")
		return "", _err
	}

	file_content := string(content)
	return strings.TrimSpace(strings.Split(file_content, "=")[1]), nil
}

func GetDefaultRegion() string {
	dirname, err := os.UserHomeDir()
	if err != nil {

		return "fr-par-1"
	}

	content, err := ioutil.ReadFile(dirname + "/.cwc/config")
	if err != nil {
		return "fr-par-1"
	}

	file_content := string(content)
	return strings.TrimSpace(strings.Split(file_content, "=")[1])
}

func SetDefaultRegion(region string) {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)

	}
	if _, err := os.Stat(dirname + "/.cwc"); os.IsNotExist(err) {
		err := os.Mkdir(dirname+"/.cwc", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := os.Create(dirname + "/.cwc/config")
	if err != nil {
		log.Fatal(err)

	}
	_, err = f.WriteString("region = " + region)
	if err != nil {
		log.Fatal(err)

	}
}
