package config

import (
	"cwc/env"
	"cwc/utils"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

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

func GetDefaultRegion() string {
	dirname, err := os.UserHomeDir()
	if nil != err {
		return "fr-par"
	}

	config_path := fmt.Sprintf("%s/.cwc/config", dirname)
	content, err := os.ReadFile(config_path)
	if nil != err {
		return "fr-par"
	}

	file_content := string(content)
	region := GetValueFromFile(file_content, "region")

	return region
}

func GetDefaultProvider() string {
	dirname, err := os.UserHomeDir()
	if nil != err {
		return ""
	}

	config_path := fmt.Sprintf("%s/.cwc/config", dirname)
	content, err := os.ReadFile(config_path)
	if nil != err {
		return ""
	}

	file_content := string(content)
	provider := GetValueFromFile(file_content, "provider")
	return provider
}

func IsPrettyFormatExpected(pretty *bool) bool {
	return *pretty || GetDefaultFormat() == "pretty"
}

func GetDefaultFormat() string {
	dirname, err := os.UserHomeDir()
	if nil != err {
		return ""
	}

	config_path := fmt.Sprintf("%s/.cwc/config", dirname)
	content, err := os.ReadFile(config_path)
	if nil != err {
		return ""
	}

	file_content := string(content)
	format := GetValueFromFile(file_content, "format")
	return format
}

func GetDefaultEndpoint() string {
	dirname, err := os.UserHomeDir()
	default_endpoint := env.API_URL
	if nil != err {
		return default_endpoint
	}

	config_path := fmt.Sprintf("%s/.cwc/config", dirname)
	content, err := os.ReadFile(config_path)
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

func SetValueToKeyInFile(file string, key string, value string) {
	dirname, err := os.UserHomeDir()
	utils.ExitIfError(err)

	file_path := fmt.Sprintf("%s/.cwc/%s", dirname, file)
	file_output, err := os.ReadFile(file_path)
	utils.ExitIfError(err)

	file_content := string(file_output)
	lines := strings.Split(file_content, "\n")
	for i, line := range lines {
		if strings.Contains(line, key+" =") {
			lines[i] = fmt.Sprintf("%s = %s", key, value)
		}
	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile(file_path, []byte(output), fs.FileMode(0644))
	utils.ExitIfError(err)
}

func UpdateFileKeyValue(filename string, key string, value string) {
	dirname, err := os.UserHomeDir()
	utils.ExitIfError(err)

	cwc_path := fmt.Sprintf("%s/.cwc", dirname)
	file_path := fmt.Sprintf("%s/%s", cwc_path, filename)
	config_path := fmt.Sprintf("%s/config", cwc_path)

	if _, err := os.Stat(cwc_path); os.IsNotExist(err) {
		err := os.Mkdir(cwc_path, os.ModePerm)
		if nil != err {
			log.Fatal(err)
		}
		os.Create(file_path)
	} else {
		if _, err := os.Stat(file_path); os.IsNotExist(err) {
			os.Create(config_path)
		}
	}

	file_content, err := os.ReadFile(file_path)
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

func SetDefaultRegion(region string) {
	UpdateFileKeyValue("config", "region", region)
}

func SetDefaultFormat(format string) {
	UpdateFileKeyValue("config", "format", format)
}

func SetDefaultProvider(provider string) {
	UpdateFileKeyValue("config", "provider", provider)
}

func SetDefaultEndpoint(endpoint string) {
	UpdateFileKeyValue("config", "endpoint", endpoint)
}

func GetUserToken() string {
	dirname, err := os.UserHomeDir()
	if nil != err {
		return ""
	}

	credentials_path := fmt.Sprintf("%s/.cwc/config", dirname)
	content, err := os.ReadFile(credentials_path)
	if nil != err {
		return ""
	}

	file_content := string(content)
	secret_key := GetValueFromFile(file_content, "cwc_secret_key")
	return secret_key
}

func AddUserCredentials(access_key string, secret_key string) {
	UpdateFileKeyValue("config", "cwc_access_key", access_key)
	UpdateFileKeyValue("config", "cwc_secret_key", secret_key)
}
