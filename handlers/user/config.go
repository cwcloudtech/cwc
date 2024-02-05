package user

import (
	"cwc/client"
	"cwc/config"
	"cwc/utils"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func HandleSwitchConfigFile(configFileName *string) {
	availableFiles, err := getFilesInFolder(".cwc")
	utils.ExitIfError(err)

	found := false
	for _, fileName := range availableFiles {
		if fileName == *configFileName {
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Config file '%s' not found\n", *configFileName)
	}

	configFilePath := filepath.Join(getHomeDir(), ".cwc", *configFileName)
	HandleImportConfigFile(configFilePath)
}

func HandleImportConfigFile(configFilePath string) {
	configContent, err := readConfigFile(configFilePath)
	utils.ExitIfError(err)

	lines := strings.Split(configContent, "\n")

	configMap := make(map[string]string)

	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			configMap[key] = value
		}
	}

	endpoint, endpointExists := configMap["endpoint"]
	format, formatExists := configMap["format"]
	provider, providerExists := configMap["provider"]
	region, regionExists := configMap["region"]
	accessKey, accessKeyExists := configMap["cwc_access_key"]
	secretKey, secretKeyExists := configMap["cwc_secret_key"]

	if endpointExists && endpoint != "" {
		config.SetDefaultEndpoint(endpoint)
		fmt.Printf("Default endpoint = %v\n", endpoint)
	}

	if formatExists && format != "" {
		config.SetDefaultFormat(format)
		fmt.Printf("Default output format = %v\n", format)
	}

	if providerExists && provider != "" {
		providers, err := client.GetProviders()
		utils.ExitIfError(err)

		availableProviders := []string{}
		for _, availableProvider := range providers.Providers {
			availableProviders = append(availableProviders, availableProvider.Name)
		}

		utils.ExitIfNeeded("Invalid provider value", !utils.StringInSlice(provider, availableProviders))

		config.SetDefaultProvider(provider)
		fmt.Printf("Default provider = %v\n", provider)
	}

	if regionExists && region != "" {
		providerRegions, err := client.GetProviderRegions()
		utils.ExitIfError(err)

		availableRegions := []string{}
		for _, availableRegion := range providerRegions.Regions {
			availableRegions = append(availableRegions, availableRegion.Name)
		}

		utils.ExitIfNeeded("Invalid region", !utils.StringInSlice(region, availableRegions))

		config.SetDefaultRegion(region)
		fmt.Printf("Default region = %v\n", region)
	}

	if accessKeyExists && secretKeyExists {
		client, _ := client.NewClient()
		err := client.UserLogin(accessKey, secretKey)
		utils.ExitIfError(err)
		fmt.Println("Credentials are set successfully")
	}

	fmt.Println("Config is set successfully")
}

func HandleGetConfigFiles() {
	fileNames, err := getFilesInFolder(".cwc")
	utils.ExitIfError(err)

	println("available config files:")
	for _, fileName := range fileNames {
		println(fileName)
	}
}

func getFilesInFolder(folderName string) ([]string, error) {
	homeDir := getHomeDir()
	folderPath := filepath.Join(homeDir, folderName)

	file, err := os.Open(folderPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	names, err := file.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	return names, nil
}

func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error getting home directory: %v", err)
	}
	return home
}

func readConfigFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
