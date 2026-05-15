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
		if strings.HasPrefix(line, key+" =") {
			requested_line = lines[i]
		}
	}

	if utils.IsBlank(requested_line) {
		return ""
	}

	return strings.TrimSpace(strings.Split(requested_line, " = ")[1])
}

func GetConfigValue(key string, defaultValue string) string {
	trimedDefaultValue := strings.TrimSpace(defaultValue)
	if envValue := os.Getenv("CWC_" + strings.ToUpper(key)); utils.IsNotBlank(envValue) {
		return strings.TrimSpace(envValue)
	}

	if envValue := os.Getenv(strings.ToUpper(key)); strings.HasPrefix(key, "cwc_") && utils.IsNotBlank(envValue) {
		return strings.TrimSpace(envValue)
	}

	dirname, err := os.UserHomeDir()
	if nil != err {
		return trimedDefaultValue
	}

	config_path := fmt.Sprintf("%s/.cwc/config", dirname)
	content, err := os.ReadFile(config_path)
	if nil != err {
		return trimedDefaultValue
	}

	file_content := string(content)
	if value := GetValueFromFile(file_content, key); utils.IsNotBlank(value) {
		return value
	}

	return trimedDefaultValue
}

func GetDefaultRegion() string {
	return GetConfigValue("region", "fr-par")
}

func GetDefaultProvider() string {
	return GetConfigValue("provider", "")
}

func GetOpenAIBaseURL() string {
	return GetConfigValue("openai_base_url", "https://api.openai.com/v1")
}

func GetOpenAIAPIKey() string {
	return GetConfigValue("openai_api_key", "")
}

func GetOpenRouterBaseURL() string {
	return GetConfigValue("openrouter_base_url", "https://openrouter.ai/api/v1")
}

func GetOpenRouterAPIKey() string {
	return GetConfigValue("openrouter_api_key", "")
}

func GetDeepSeekBaseURL() string {
	return GetConfigValue("deepseek_base_url", "https://api.deepseek.com/v1")
}

func GetDeepSeekAPIKey() string {
	return GetConfigValue("deepseek_api_key", "")
}

func GetAnthropicBaseURL() string {
	return GetConfigValue("anthropic_base_url", "https://api.anthropic.com/v1")
}

func GetAnthropicAPIKey() string {
	return GetConfigValue("anthropic_api_key", "")
}

func GetGeminiBaseURL() string {
	return GetConfigValue("gemini_base_url", "https://generativelanguage.googleapis.com/v1beta")
}

func GetGeminiAPIKey() string {
	return GetConfigValue("gemini_api_key", "")
}

func GetDefaultAiProvider() string {
	return GetConfigValue("default_ai_provider", "openai")
}

func GetDefaultAiModel() string {
	return GetConfigValue("default_ai_model", "")
}

func GetGitLabBaseURL() string {
	return GetConfigValue("gitlab_base_url", "https://gitlab.com/api/v4")
}

func GetGitLabToken() string {
	return GetConfigValue("gitlab_token", "")
}

func GetGitLabWebhookSecret() string {
	return GetConfigValue("gitlab_webhook_secret", "")
}

func GetDefaultKubeConfigPath() string {
	return GetConfigValue("kube_config_path", "")
}

func GetDefaultFormat() string {
	return GetConfigValue("format", "")
}

func IsPrettyFormatExpected(pretty *bool) bool {
	return *pretty || GetDefaultFormat() == "pretty"
}

func GetDefaultEndpoint() string {
	return GetConfigValue("endpoint", env.API_URL)
}

func GetRepoURL() string {
	return GetConfigValue("repo_url", env.REPO_URL)
}

func GetRepoBranch() string {
	return GetConfigValue("repo_branch", env.BRANCH)
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
		if strings.HasPrefix(line, key+" =") {
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

func SetDefaultKubeConfigPath(path string) {
	UpdateFileKeyValue("config", "kube_config_path", path)
}

func GetUserToken() string {
	return GetConfigValue("cwc_secret_key", "")
}

func AddUserCredentials(access_key string, secret_key string) {
	UpdateFileKeyValue("config", "cwc_access_key", access_key)
	UpdateFileKeyValue("config", "cwc_secret_key", secret_key)
}
