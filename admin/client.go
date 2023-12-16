package admin

import (
	"bytes"
	"cwc/env"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func NewClient() (*Client, error) {
	region := GetDefaultRegion()
	provider := GetDefaultProvider()
	err := error(nil)

	if provider == "" {
		err = fmt.Errorf("default provider is not set")
	} else if region == "" {
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
		if errorResponse.Error == "" {
			return nil, fmt.Errorf(fmt.Sprintf("Request failed with status %d", resp.StatusCode))
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

	if nil != err {
		log.Fatal(err)
	}

	if _, err := os.Stat(dirname + "/.cwc"); os.IsNotExist(err) {
		err := os.Mkdir(dirname+"/.cwc", os.ModePerm)
		if nil != err {
			log.Fatal(err)
		}
	}

	f, err := os.Create(dirname + "/.cwc/credentials")
	if nil != err {
		log.Fatal(err)
	}

	_, err = f.WriteString("cwc_access_key = " + access_key + "\n")
	if nil != err {
		log.Fatal(err)
	}

	_, err = f.WriteString("cwc_secret_key = " + secret_key + "\n")
	if nil != err {
		log.Fatal(err)
	}
}

func getUserToken() string {
	dirname, err := os.UserHomeDir()
	if nil != err {
		return ""
	}

	content, err := ioutil.ReadFile(dirname + "/.cwc/credentials")
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

	content, err := ioutil.ReadFile(dirname + "/.cwc/config")
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

	content, err := ioutil.ReadFile(dirname + "/.cwc/config")
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

	content, err := ioutil.ReadFile(dirname + "/.cwc/config")
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

	content, err := ioutil.ReadFile(dirname + "/.cwc/config")
	if nil != err {
		return default_endpoint
	}

	file_content := string(content)
	endpoint := GetValueFromFile(file_content, "endpoint")
	if "" == endpoint {
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

	if requested_line == "" {
		return ""
	}

	return strings.Split(requested_line, " = ")[1]
}

func UpdateFileKeyValue(filename string, key string, value string) {
	dirname, err := os.UserHomeDir()

	if nil != err {
		log.Fatal(err)
	}

	if _, err := os.Stat(dirname + "/.cwc"); os.IsNotExist(err) {
		err := os.Mkdir(dirname+"/.cwc", os.ModePerm)
		if nil != err {
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
		if nil != err {
			log.Fatal(err)
		}

		_, err = config_file.WriteString(key + " = " + value + "\n")
		if nil != err {
			log.Fatal(err)

		}
	} else {
		SetValueToKeyInFile(filename, key, value)
	}

}
func SetValueToKeyInFile(file string, key string, value string) {
	dirname, err := os.UserHomeDir()
	if nil != err {
		log.Fatalln(err)
	}

	file_output, err := ioutil.ReadFile(dirname + "/.cwc/" + file)
	if nil != err {
		log.Fatalln(err)
	}

	file_content := string(file_output)
	lines := strings.Split(file_content, "\n")
	for i, line := range lines {
		if strings.Contains(line, key+" =") {
			lines[i] = key + " = " + value
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(dirname+"/.cwc/"+file, []byte(output), 0644)
	if nil != err {
		log.Fatalln(err)
	}
}
