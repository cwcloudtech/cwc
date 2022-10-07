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

func NewClient() *Client {
	region := GetDefaultRegion()
	provider := GetDefaultProvider()
	return &Client{
		region:     region,
		provider:   provider,
		httpClient: &http.Client{},
	}
}

func (c *Client) UserLogin(access_key string, secret_key string) error {
	addUserCredentials(access_key, secret_key)
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
		if errorResponse.Error == "" {
			return nil, fmt.Errorf(fmt.Sprintf("Request failed with status %d", resp.StatusCode))

		} else {

			return nil, fmt.Errorf(errorResponse.Error)
		}

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
	_, err = f.WriteString("cwc_access_key = " + access_key + "\n")

	if err != nil {
		log.Fatal(err)

	}
	_, err = f.WriteString("cwc_secret_key = " + secret_key + "\n")

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
	secret_key := GetValueFromFile(file_content, "cwc_secret_key")
	return secret_key, err
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
	region := GetValueFromFile(file_content, "region")
	return region
}

func SetDefaultRegion(region string) {
	UpdateFileKeyValue("config", "region", region)
}

func SetDefaultProvider(provider string) {
	UpdateFileKeyValue("config", "provider", provider)
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
	provider := GetValueFromFile(file_content, "provider")
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
	endpoint := GetValueFromFile(file_content, "endpoint")
	return endpoint
}

func SetDefaultEndpoint(endpoint string) {
	UpdateFileKeyValue("config", "endpoint", endpoint)
}

func GetValueFromFile(content_file string, key string) string {
	lines := strings.Split(content_file, "\n")
	var requested_line string
	for i, line := range lines {
		if strings.Contains(line, key+" =") {
			requested_line = lines[i]
		}
	}
	if requested_line == "" {
		return ""
	}
	return strings.Split(requested_line, " = ")[1]
}

func UpdateFileKeyValue(filename string, key string, value string) {
	dirname, err := os.UserHomeDir()

	if err != nil {
		log.Fatal(err)

	}
	if _, err := os.Stat(dirname + "/.cwc"); os.IsNotExist(err) {
		err := os.Mkdir(dirname+"/.cwc", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		os.Create(dirname + "/.cwc/" + filename)
	} else {
		if _, err := os.Stat(dirname + "/.cwc/" + filename); os.IsNotExist(err) {
			os.Create(dirname + "/.cwc/config")
		}
	}
	file_content, err := ioutil.ReadFile(dirname + "/.cwc/" + filename)
	if GetValueFromFile(string(file_content), key) == "" {
		config_file, err := os.OpenFile(dirname+"/.cwc/"+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		_, err = config_file.WriteString(key + " = " + value + "\n")
		if err != nil {
			log.Fatal(err)

		}
	} else {
		SetValueToKeyInFile(filename, key, value)
	}

}
func SetValueToKeyInFile(file string, key string, value string) {
	dirname, err := os.UserHomeDir()
	file_output, err := ioutil.ReadFile(dirname + "/.cwc/" + file)
	file_content := string(file_output)
	lines := strings.Split(file_content, "\n")
	for i, line := range lines {
		if strings.Contains(line, key+" =") {
			lines[i] = key + " = " + value + "\n"
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(dirname+"/.cwc/"+file, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

}
