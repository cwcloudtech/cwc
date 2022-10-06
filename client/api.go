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

type AttachInstanceRequest struct {
	ProjectId     int    `json:"project_id"`
	Name          string `json:"name"`
	Instance_type string `json:"type"`
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
	Instances []Instance `json:"instances"`
}

type AddProjectBody struct {
	Name      string     `json:"name"`
	Host      string     `json:"host"`
	Token      string     `json:"token"`
	Namespace      string     `json:"namespace"`
	GitUsername      string     `json:"git_username"`

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

func (c *Client) AddProject(project_name string,host string,token string,git_username string,namespace string) (*Project, error) {
	buf := bytes.Buffer{}
	project := AddProjectBody{
		Name:   project_name,
		Host: host,
		Token: token,
		GitUsername: git_username,
		Namespace: namespace,
	}

	err := json.NewEncoder(&buf).Encode(project)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/project"), "POST", buf)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	created_project := &Project{}
	err = json.NewDecoder(respBody).Decode(created_project)
	if err != nil {
		return nil, err
	}
	return created_project, nil
}

func (c *Client) DeleteProject(projectId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/project/%s", projectId), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetAllProjects() (*[]Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/project"), "GET", bytes.Buffer{})
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	projects := []Project{}
	err = json.NewDecoder(body).Decode(&projects)

	if err != nil {
		fmt.Printf("errr")
		return nil, err
	}
	return &projects, nil
}

func (c *Client) GetAllEnvironments() (*[]Environment, error) {
	body, err := c.httpRequest(fmt.Sprintf("/environment/all"), "GET", bytes.Buffer{})
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
func (c *Client) UserLogin(access_key string, secret_key string) error {
	addUserCredentials(access_key,secret_key)
	return nil
}
func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)

	user_token, err := getUserToken()
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Auth-Token", user_token)

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
	hostname := GetDefaultEndpoint()
	return fmt.Sprintf("%s%s", hostname, path)
}

func addUserCredentials(access_key string, secret_key string) {
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
	_, err = f.WriteString("cwc_access_key = " + access_key+"\n")

	if err != nil {
		log.Fatal(err)

	}
	_, err = f.WriteString("cwc_secret_key = " + secret_key+"\n")

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
	secret_key := GetValueFromFile(file_content,"cwc_secret_key" )
	return secret_key,err
}

func GetDefaultRegion() string {
	dirname, err := os.UserHomeDir()
	if err != nil {

		return "fr-par"
	}

	content, err := ioutil.ReadFile(dirname + "/.cwc/config")
	if err != nil {
		return "fr-par"
	}

	file_content := string(content)
	region := GetValueFromFile(file_content,"region")
	return region
}

func SetDefaultRegion(region string) {
	UpdateFileKeyValue("config","region",region)
}

func SetDefaultProvider(provider string) {
	UpdateFileKeyValue("config","provider",provider)
}


func GetDefaultProvider() string {
	dirname, err := os.UserHomeDir()
	if err != nil {

		return "None"
	}

	content, err := ioutil.ReadFile(dirname + "/.cwc/config")
	if err != nil {
		return "None"
	}

	file_content := string(content)
	provider := GetValueFromFile(file_content,"provider")
	return provider
}


func GetDefaultEndpoint() string {
	dirname, err := os.UserHomeDir()
	if err != nil {

		return "None"
	}

	content, err := ioutil.ReadFile(dirname + "/.cwc/config")
	if err != nil {
		return "None"
	}

	file_content := string(content)
	endpoint := GetValueFromFile(file_content,"endpoint")
	return endpoint
}



func SetDefaultEndpoint(endpoint string) {
	UpdateFileKeyValue("config","endpoint",endpoint)
}

func GetValueFromFile(content_file string,key string) string {
	lines := strings.Split(content_file, "\n")
	var requested_line string
	for i, line := range lines {
			if strings.Contains(line, key+" =") {
				requested_line = lines[i]
			}
	}
	if requested_line ==""{
		return ""
	}
	return strings.Split(requested_line, " = ")[1]
}


func UpdateFileKeyValue(filename string, key string, value string){
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)

	}
	if _, err := os.Stat(dirname + "/.cwc"); os.IsNotExist(err) {
		err := os.Mkdir(dirname+"/.cwc", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		os.Create(dirname + "/.cwc/"+filename)
	}else{
		if _, err := os.Stat(dirname+"/.cwc/"+filename);os.IsNotExist(err) {
			os.Create(dirname + "/.cwc/config")
		  }
	}
	file_content, err := ioutil.ReadFile(dirname + "/.cwc/"+filename)
	if GetValueFromFile(string(file_content),key)==""{
		config_file, err := os.OpenFile(dirname + "/.cwc/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		_, err = config_file.WriteString(key+" = " + value+"\n")
		if err != nil {
			log.Fatal(err)
	
		}
	}else{
		SetValueToKeyInFile(filename,key,value)
	}

}
func SetValueToKeyInFile(file string,key string,value string) {
	dirname, err := os.UserHomeDir()
	file_output, err := ioutil.ReadFile(dirname + "/.cwc/"+file)
	file_content := string(file_output)
	lines := strings.Split(file_content, "\n")
	for i, line := range lines {
			if strings.Contains(line, key+" =") {
				lines[i] = key+" = "+value+"\n"
			}
	}
	output := strings.Join(lines, "\n")
    err = ioutil.WriteFile(dirname + "/.cwc/"+file, []byte(output), 0644)
    if err != nil {
                log.Fatalln(err)
    }

}