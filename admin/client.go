package admin

import (
	"bytes"
	"cwc/env"
	"cwc/utils"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func NewClient() (*Client, error) {
	region := GetDefaultRegion()
	provider := GetDefaultProvider()
	err := error(nil)

	if utils.IsBlank(provider) {
		err = fmt.Errorf("default provider is not set")
	} else if utils.IsBlank(region) {
		err = fmt.Errorf("default region is not set")
	}

	return &Client{
		region:     region,
		provider:   provider,
		httpClient: &http.Client{},
	}, err
}

func (c *Client) UserLogin(access_key string, secret_key string) error {

	buf := bytes.Buffer{}
	project := ApiKey{
		Accesskey: access_key,
		SecretKey: secret_key,
	}

	err := json.NewEncoder(&buf).Encode(project)
	if nil != err {
		return err
	}
	_, err = c.httpRequest("/api_keys/verify", "POST", buf)
	if nil != err {
		return err
	}

	addUserCredentials(access_key, secret_key)
	return nil
}
func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)
	if nil != err {
		return nil, err
	}

	user_token := getUserToken()

	req.Header.Set("X-Auth-Token", user_token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if nil != err {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if nil != err {
			return nil, fmt.Errorf("an error occured")
		}

		errorResponse := ErrorResponse{}
		json.NewDecoder(respBody).Decode(&errorResponse)
		if utils.IsBlank(errorResponse.Error) {
			return nil, fmt.Errorf("request failed with status %d", resp.StatusCode)
		} else {
			return nil, fmt.Errorf(errorResponse.Error)
		}
	}

	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	default_api_version := "v1"
	hostname := GetDefaultEndpoint()
	return fmt.Sprintf("%s/%s%s", hostname, default_api_version, path)
}

func addUserCredentials(access_key string, secret_key string) {
	dirname, err := os.UserHomeDir()
	utils.ExitIfError(err)

	cwc_path := fmt.Sprintf("%s/.cwc", dirname)
	if _, err := os.Stat(cwc_path); os.IsNotExist(err) {
		err := os.Mkdir(cwc_path, os.ModePerm)
		utils.ExitIfError(err)
	}

	credentials_path := fmt.Sprintf("%s/credentials", cwc_path)
	f, err := os.Create(credentials_path)
	utils.ExitIfError(err)

	_, err = f.WriteString(fmt.Sprintf("cwc_access_key = %s\n", access_key))
	utils.ExitIfError(err)

	_, err = f.WriteString(fmt.Sprintf("cwc_secret_key = %s\n", secret_key))
	utils.ExitIfError(err)
}

func getUserToken() string {
	dirname, err := os.UserHomeDir()
	if nil != err {
		return ""
	}

	credentials_path := fmt.Sprintf("%s/.cwc/credentials", dirname)
	content, err := ioutil.ReadFile(credentials_path)
	if nil != err {
		return ""
	}

	file_content := string(content)
	secret_key := GetValueFromFile(file_content, "cwc_secret_key")
	return secret_key
}

func GetDefaultRegion() string {
	dirname, err := os.UserHomeDir()
	if nil != err {
		return "fr-par"
	}

	configDir := fmt.Sprintf("%s/.cwc/config", dirname)
	content, err := ioutil.ReadFile(configDir)
	if nil != err {
		return "fr-par"
	}

	file_content := string(content)
	region := GetValueFromFile(file_content, "region")
	return region
}

func SetDefaultRegion(region string) {
	UpdateFileKeyValue("config", "region", region)
}

func SetDefaultFormat(format string) {
	UpdateFileKeyValue("config", "format", format)
}

func SetDefaultProvider(provider string) {
	UpdateFileKeyValue("config", "provider", provider)
}

func GetDefaultFormat() string {
	dirname, err := os.UserHomeDir()
	if nil != err {
		return ""
	}

	configDir := fmt.Sprintf("%s/.cwc/config", dirname)
	content, err := ioutil.ReadFile(configDir)
	if nil != err {
		return ""
	}

	file_content := string(content)
	format := GetValueFromFile(file_content, "format")
	return format
}

func GetDefaultProvider() string {
	dirname, err := os.UserHomeDir()
	if nil != err {
		return ""
	}

	configDir := fmt.Sprintf("%s/.cwc/config", dirname)
	content, err := ioutil.ReadFile(configDir)
	if nil != err {
		return ""
	}

	file_content := string(content)
	provider := GetValueFromFile(file_content, "provider")
	return provider
}

func GetDefaultEndpoint() string {
	dirname, err := os.UserHomeDir()
	default_endpoint := env.API_URL
	if nil != err {
		return default_endpoint
	}

	configDir := fmt.Sprintf("%s/.cwc/config", dirname)
	content, err := ioutil.ReadFile(configDir)
	if nil != err {
		return default_endpoint
	}

	file_content := string(content)
	endpoint := GetValueFromFile(file_content, "endpoint")
	if utils.IsBlank(endpoint) {
		return default_endpoint
	}

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

	if utils.IsBlank(requested_line) {
		return ""
	}

	return strings.Split(requested_line, " = ")[1]
}

func UpdateFileKeyValue(filename string, key string, value string) {
	dirname, err := os.UserHomeDir()
	utils.ExitIfError(err)

	dir_path := fmt.Sprintf("%s/.cwc", dirname)
	file_path := fmt.Sprintf("%s/%s", dir_path, filename)
	config_path := fmt.Sprintf("%s/config", file_path)

	if _, err := os.Stat(dir_path); os.IsNotExist(err) {
		err := os.Mkdir(dir_path, os.ModePerm)
		utils.ExitIfError(err)

		os.Create(file_path)
	} else {
		if _, err := os.Stat(file_path); os.IsNotExist(err) {
			os.Create(config_path)
		}
	}

	file_content, err := ioutil.ReadFile(file_path)
	utils.ExitIfError(err)

	if utils.IsBlank(GetValueFromFile(string(file_content), key)) {
		config_file, err := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		utils.ExitIfError(err)

		_, err = config_file.WriteString(fmt.Sprintf("%s = %s\n", key, value))
		utils.ExitIfError(err)
	} else {
		SetValueToKeyInFile(filename, key, value)
	}

}
func SetValueToKeyInFile(file string, key string, value string) {
	dirname, err := os.UserHomeDir()
	utils.ExitIfError(err)

	file_path := fmt.Sprintf("%s/.cwc/%s", dirname, file)
	file_output, err := ioutil.ReadFile(file_path)
	utils.ExitIfError(err)

	file_content := string(file_output)
	lines := strings.Split(file_content, "\n")
	for i, line := range lines {
		if strings.Contains(line, key+" =") {
			lines[i] = fmt.Sprintf("%s = %s", key, value)
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(file_path, []byte(output), 0644)
	utils.ExitIfError(err)
}
