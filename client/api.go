package client

import (
	"bytes"
	"encoding/json"
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

type UpdateProjectRequest struct {
	Status        string `json:"status"`
	Instance_type string `json:"type"`
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

type Project struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Gitlab_url    string `json:"gitlab_project_url"`
	Instance_type string `json:"type"`
	Environment   string `json:"environment"`
	Status        string `json:"status"`
	Email         string `json:"email"`
	Region        string `json:"region"`
	Ip_address    string `json:"ip_address"`
}

func NewClient(region string) *Client {
	return &Client{
		region:     region,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetAll() (*[]Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/instance/%s", c.region), "GET", bytes.Buffer{})
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

func (c *Client) GetProject(project_id string) (*Project, error) {
	body, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s", c.region, project_id), "GET", bytes.Buffer{})
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

func (c *Client) AddProject(project_name string, instance_size string, environment string, email string) (*Project, error) {
	buf := bytes.Buffer{}
	project := Project{
		Name:          project_name,
		Instance_type: instance_size,
		Environment:   environment,
		Email:         email,
		Region:        c.region,
	}

	err := json.NewEncoder(&buf).Encode(project)
	if err != nil {
		return nil, err
	}
	respBody, err := c.httpRequest(fmt.Sprintf("/instance/%s/provision/%s", c.region, project.Environment), "POST", buf)
	if err != nil {
		return nil, err
	}
	created_project := &Project{}
	err = json.NewDecoder(respBody).Decode(created_project)
	if err != nil {
		return nil, err
	}
	return created_project, nil
}

func (c *Client) UpdateProject(id string, status string, instance_type string) error {
	buf := bytes.Buffer{}

	updateProjectRequest := &UpdateProjectRequest{
		Status:        status,
		Instance_type: instance_type,
	}
	err := json.NewEncoder(&buf).Encode(updateProjectRequest)
	if err != nil {
		return err
	}
	_, err = c.httpRequest(fmt.Sprintf("/instance/%s/%s", c.region, id), "PATCH", buf)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteProject(projectId string) error {
	_, err := c.httpRequest(fmt.Sprintf("/instance/%s/%s", c.region, projectId), "DELETE", bytes.Buffer{})
	if err != nil {
		return err
	}
	return nil
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
	if resp.StatusCode != http.StatusOK {
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
		fmt.Printf("cwc: access denied, please login\n")
		return "", err
	}

	content, err := ioutil.ReadFile(dirname + "/.cwc/credentials")
	if err != nil {
		fmt.Printf("cwc: access denied, please login\n")
		return "", err
	}

	file_content := string(content)
	return strings.TrimSpace(strings.Split(file_content, "=")[1]), nil
}
