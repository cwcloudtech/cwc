package user

import (
	"crypto/rand"
	"encoding/base64"
	"cwc/config"
	"cwc/env"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"cwc/utils"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var keyValuePattern = regexp.MustCompile(`^[^=]+=[^=]+$`)
var yamlFilePattern = regexp.MustCompile(`(?i)^.+\.ya?ml$`)

type RepoConfig struct {
	RepoURL  string
	Branch   string
	Username string
	Password string
}

type PortForwardConfig struct {
	Name       string `yaml:"name"`
	Port       int    `yaml:"port"`
	TargetPort int    `yaml:"targetPort"`
}

type PfwConfigFile struct {
	PfwConfigs []PortForwardConfig `yaml:"pfwConfigs"`
}

func GetRepoConfig() RepoConfig {
	repoURL := config.GetRepoURL()
	branch := config.GetRepoBranch()

	if utils.IsBlank(repoURL) {
		repoURL = env.REPO_URL
	} else {
		env.REPO_URL = repoURL
	}

	if utils.IsBlank(branch) {
		branch = env.BRANCH
	} else {
		env.BRANCH = branch
	}

	return RepoConfig{
		RepoURL: repoURL,
		Branch:  branch,
	}
}

func (c *RepoConfig) SetRepoURL(url string) {
	config.UpdateFileKeyValue("config", "repo_url", url)
	env.REPO_URL = url
}

func (c *RepoConfig) SetRepoBranch(branch string) {
	config.UpdateFileKeyValue("config", "repo_branch", branch)
	env.BRANCH = branch
}

func SaveRepoConfig(config *RepoConfig) {
	if utils.IsNotBlank(config.RepoURL) {
		config.SetRepoURL(config.RepoURL)
	}

	if utils.IsNotBlank(config.Branch) {
		config.SetRepoBranch(config.Branch)
	}

	if utils.IsNotBlank(config.Username) {
		env.REPO_USERNAME = config.Username
	}

	if utils.IsNotBlank(config.Password) {
		env.REPO_PASSWORD = config.Password
	}
}

func HandleBootstrap(cmd *cobra.Command, releaseName, nameSpace string, kindCluster string, otherValues []string, flagVerbose bool, keepDir bool, recreateNs bool, openshift bool) {
	config := GetRepoConfig()
	repoURL := config.RepoURL
	directory := env.DIRECTORY
	branch := config.Branch

	if err := CloneRepo(repoURL, directory, branch, keepDir, env.REPO_USERNAME, env.REPO_PASSWORD); err != nil {
		log.Printf("Error cloning repository: %v", err)
		return
	}

	password, err := replaceVariablesInValues(directory, nameSpace)
	if err != nil || utils.IsBlank(password) {
		log.Printf("Error updating values.yaml: %v", err)
		return
	} else {
		log.Printf("Generated random password for tenant: %s", password)
	}

	log.Println("Starting Helm chart installation...")

	if err := runKindRecreateCluster(kindCluster); err != nil {
		log.Fatalf("Error running kind command: %v", err)
	}

	if err := runHelmDependancyUpdate(directory, keepDir); err != nil {
		log.Fatalf("Error running helm command: %v", err)
	}

	if err := runDeleteNS(nameSpace, recreateNs, openshift); err != nil {
		log.Printf("Not able to delete the namespace: %s, error: %v", nameSpace, err)
	}

	if err := runCreateNS(nameSpace, openshift); err != nil {
		log.Printf("Not able to create the namespace: %s, error: %v", nameSpace, err)
	}

	if err := runHelmInstall(releaseName, directory, nameSpace, otherValues, openshift); err != nil {
		log.Fatalf("Error running helm command: %v", err)
	}

	if err := runGitClean(directory, config); err != nil {
		log.Printf("Error running git clean commands: %v", err)
	}

	log.Println("Helm chart installation started in background.")
}

func runDeleteNS(nameSpace string, recreateNs bool, openshift bool) error {
	if !recreateNs {
		return nil
	}

	kubectlCommand := utils.If(openshift, "oc", "kubectl")

	kubectlArgs := []string{
		"delete",
		"ns",
		nameSpace,
	}

	log.Printf("Executing %s command: %s %s", kubectlCommand, kubectlCommand, strings.Join(kubectlArgs, " "))

	kubectlDeleteNs := exec.Command(kubectlCommand, kubectlArgs...)
	kubectlDeleteNs.Stdout = os.Stdout
	kubectlDeleteNs.Stderr = os.Stderr

	return kubectlDeleteNs.Run()
}

func runCreateNS(nameSpace string, openshift bool) error {
	kubectlCommand := utils.If(openshift, "oc", "kubectl")

	kubectlArgs := []string{
		"create",
		"ns",
		nameSpace,
	}

	log.Printf("Executing %s command: %s %s", kubectlCommand, kubectlCommand, strings.Join(kubectlArgs, " "))

	kubectlDeleteNs := exec.Command(kubectlCommand, kubectlArgs...)
	kubectlDeleteNs.Stdout = os.Stdout
	kubectlDeleteNs.Stderr = os.Stderr

	return kubectlDeleteNs.Run()
}

func runHelmDependancyUpdate(directory string, keepDir bool) error {
	if _, err := os.Stat(directory + "/charts"); !os.IsNotExist(err) {
		if keepDir {
			return nil
		}
	}

	helmCommand := "helm"
	helmArgs := []string{
		"dependency",
		"update",
	}

	log.Printf("Executing helm command: %s %s", helmCommand, strings.Join(helmArgs, " "))

	helmDepUdpate := exec.Command(helmCommand, helmArgs...)
	helmDepUdpate.Dir = directory
	helmDepUdpate.Stdout = os.Stdout
	helmDepUdpate.Stderr = os.Stderr

	return helmDepUdpate.Run()
}

func runKindRecreateCluster(clusterName string) error {
	if utils.IsBlank(clusterName) {
		return nil
	}

	kindCommand := "kind"
	kindArgs := []string{
		"delete",
		"cluster",
		"--name", clusterName,
	}

	log.Printf("Executing kind command: %s %s", kindCommand, strings.Join(kindArgs, " "))

	kindDeleteCluster := exec.Command(kindCommand, kindArgs...)
	kindDeleteCluster.Stdout = os.Stdout
	kindDeleteCluster.Stderr = os.Stderr

	if err := kindDeleteCluster.Run(); err != nil {
		log.Printf("error deleting kind cluster: %v", err)
	}

	kindArgs = []string{
		"create",
		"cluster",
		"--name", clusterName,
	}

	log.Printf("Executing kind command: %s %s", kindCommand, strings.Join(kindArgs, " "))

	kindCreateCluster := exec.Command(kindCommand, kindArgs...)
	kindCreateCluster.Stdout = os.Stdout
	kindCreateCluster.Stderr = os.Stderr

	return kindCreateCluster.Run()
}

func runHelmInstall(releaseName string, directory string, nameSpace string, otherValues []string, openshift bool) error {
	helmCommand := "helm"
	helmArgs := utils.If(openshift, []string{
		"install",
		releaseName,
		directory,
		"--namespace", nameSpace,
		"--set", "s3.enabled=false",
	}, []string{
		"install",
		releaseName,
		directory,
		"--namespace", nameSpace,
	})

	for _, value := range otherValues {
		trimmedValue := strings.TrimSpace(value)

		switch {
		case keyValuePattern.MatchString(trimmedValue):
			helmArgs = append(helmArgs, "--set", trimmedValue)
		case yamlFilePattern.MatchString(trimmedValue):
			helmArgs = append(helmArgs, "-f", trimmedValue)
		default:
			return fmt.Errorf("invalid --value %q: expected key=value or .yaml/.yml file", value)
		}
	}

	log.Printf("Executing helm command: %s %s", helmCommand, strings.Join(helmArgs, " "))

	helmInstallation := exec.Command(helmCommand, helmArgs...)
	helmInstallation.Stdout = os.Stdout
	helmInstallation.Stderr = os.Stderr
	helmInstallation.SysProcAttr = helmInstallSysProcAttr()

	if err := helmInstallation.Start(); err != nil {
		return err
	}

	log.Printf("Helm install is running in background with PID %d", helmInstallation.Process.Pid)

	return helmInstallation.Process.Release()
}

func CloneRepo(repoURL, directory, branch string, keepDir bool, username, password string) error {
	if _, err := os.Stat(directory); !os.IsNotExist(err) {
		if keepDir {
			return nil
		}

		fmt.Printf("Deleting existing directory: %s\n", directory)
		if err := os.RemoveAll(directory); err != nil {
			return fmt.Errorf("failed to delete existing directory: %v", err)
		}
	}

	cloneOptions := &git.CloneOptions{
		URL:           repoURL,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		Progress:      os.Stdout,
	}

	if utils.IsNotBlank(username) || utils.IsNotBlank(password) {
		cloneOptions.Auth = &http.BasicAuth{
			Username: username,
			Password: password,
		}
	}

	_, err := git.PlainClone(directory, false, cloneOptions)

	if err != nil {
		return fmt.Errorf("failed to clone repository: %v", err)
	}

	fmt.Println("Repository cloned successfully.")
	return nil
}

func runGitClean(directory string, config RepoConfig) error {
	gitCommand := "git"

	commandArgs := [][]string{{"add", "."}, {"stash"}, {"stash", "clear"}, {"reset", "--hard", "origin/" + config.Branch}}

	for _, args := range commandArgs {
		log.Printf("Executing git command: %s %s", gitCommand, strings.Join(args, " "))
		gitCmd := exec.Command(gitCommand, args...)
		gitCmd.Dir = directory
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		if err := gitCmd.Run(); err != nil {
			return fmt.Errorf("failed to execute git command %v: %w", args, err)
		}
	}

	return nil
}

func generateRandomPassword(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("password length must be greater than 0")
	}

	byteLen := (length*3 + 3) / 4
	randomBytes := make([]byte, byteLen)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	return base64.RawURLEncoding.EncodeToString(randomBytes)[:length], nil
}

func replaceVariablesInValues(chartDirectory string, nameSpace string) (string, error) {
	patterns := []string{
		"values.yaml",
		"values.yml",
		"values-*.yaml",
		"values-*.yml",
		"values.*.yaml",
		"values.*.yml",
	}

	valuesFilePaths := make(map[string]struct{})
	for _, pattern := range patterns {
		matchedPaths, err := filepath.Glob(filepath.Join(chartDirectory, pattern))
		if err != nil {
			return "", fmt.Errorf("invalid values file pattern %q: %w", pattern, err)
		}

		for _, matchedPath := range matchedPaths {
			fileInfo, err := os.Stat(matchedPath)
			if err != nil {
				return "", fmt.Errorf("failed to stat values file %q: %w", matchedPath, err)
			}

			if !fileInfo.Mode().IsRegular() {
				continue
			}

			valuesFilePaths[matchedPath] = struct{}{}
		}
	}

	if len(valuesFilePaths) == 0 {
		return "", fmt.Errorf("no values files found in %q", chartDirectory)
	}

	password, err := generateRandomPassword(20)
	if err != nil {
		return "", fmt.Errorf("failed to generate password: %w", err)
	}

	for valuesFilePath := range valuesFilePaths {
		content, err := os.ReadFile(valuesFilePath)
		if err != nil {
			return "", fmt.Errorf("failed to read values file %q: %w", valuesFilePath, err)
		}

		replacedContent := strings.ReplaceAll(string(content), "${{tenant-name}}", nameSpace)
		replacedContent = strings.ReplaceAll(replacedContent, "${{namespace}}", nameSpace)
		replacedContent = strings.ReplaceAll(replacedContent, "${{tenant-password}}", password)
		replacedContent = strings.ReplaceAll(replacedContent, "${{password}}", password)
		if replacedContent == string(content) {
			continue
		}

		fileInfo, err := os.Stat(valuesFilePath)
		if err != nil {
			return password, fmt.Errorf("failed to stat values file %q: %w", valuesFilePath, err)
		}

		if err := os.WriteFile(valuesFilePath, []byte(replacedContent), fileInfo.Mode()); err != nil {
			return password, fmt.Errorf("failed to update values file %q: %w", valuesFilePath, err)
		}
	}

	return password, nil
}

func HandleUninstall(cmd *cobra.Command, releaseName string, nameSpace string, force bool, openshift bool) {
	log.Println("Starting Helm chart uninstallation...")

	if err := runHelmUninstall(releaseName, nameSpace); err != nil {
		if !force {
			log.Fatalf("Error running helm uninstall command: %v", err)
		} else {
			log.Printf("Error running helm uninstall command: %v", err)
		}
	}

	if err := runDeleteAll(nameSpace, force, openshift); err != nil {
		log.Printf("Error running kubectl delete all command: %v", err)
	}

	log.Println("Helm chart uninstallation completed successfully.")
}

func runHelmUninstall(releaseName, nameSpace string) error {
	helmCommand := "helm"
	helmArgs := []string{
		"uninstall",
		releaseName,
		"--namespace", nameSpace,
	}

	log.Printf("Executing helm command: %s %s", helmCommand, strings.Join(helmArgs, " "))

	helmUninstallation := exec.Command(helmCommand, helmArgs...)
	helmUninstallation.Stdout = os.Stdout
	helmUninstallation.Stderr = os.Stderr

	return helmUninstallation.Run()
}

func runDeleteAll(nameSpace string, force bool, openshift bool) error {
	if !force {
		return nil
	}

	kubectlCommand := utils.If(openshift, "oc", "kubectl")

	kubectlArgs := []string{
		"-n",
		nameSpace,
		"delete",
		"all",
		"--all",
	}

	log.Printf("Executing %s command: %s %s", kubectlCommand, kubectlCommand, strings.Join(kubectlArgs, " "))

	kubectlDeleteAll := exec.Command(kubectlCommand, kubectlArgs...)
	kubectlDeleteAll.Stdout = os.Stdout
	kubectlDeleteAll.Stderr = os.Stderr

	return kubectlDeleteAll.Run()
}

func HandlePortForward(cmd *cobra.Command, nameSpace string, openshift bool, configPath string) {
	log.Println("Starting tunnels...")

	portForwardConfigs, err := loadPortForwardConfigs(configPath)
	if err != nil {
		log.Fatalf("Error loading pfw config %q: %v", configPath, err)
	}

	log.Printf("Loaded %d port-forward entries from %s", len(portForwardConfigs), configPath)

	for _, cfg := range portForwardConfigs {
		if err := runPortForward(nameSpace, cfg.Name, cfg.Port, cfg.TargetPort, openshift); err != nil {
			log.Fatalf("Error running kubectl: %v", err)
		}

		log.Printf("Going to %s: http://localhost:%d", cfg.Name, cfg.TargetPort)
	}
}

func loadPortForwardConfigs(configPath string) ([]PortForwardConfig, error) {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config PfwConfigFile
	if err := yaml.Unmarshal(content, &config); err != nil {
		return nil, err
	}

	if len(config.PfwConfigs) == 0 {
		return nil, fmt.Errorf("no pfwConfigs entries found")
	}

	for i, cfg := range config.PfwConfigs {
		if utils.IsBlank(cfg.Name) {
			return nil, fmt.Errorf("pfwConfigs[%d].name is required", i)
		}
		if cfg.Port <= 0 {
			return nil, fmt.Errorf("pfwConfigs[%d].port must be greater than 0", i)
		}
		if cfg.TargetPort <= 0 {
			return nil, fmt.Errorf("pfwConfigs[%d].targetPort must be greater than 0", i)
		}
	}

	return config.PfwConfigs, nil
}

func runPortForward(nameSpace string, service string, port int, targetPort int, openshift bool) error {
	kubectlCommand := utils.If(openshift, "oc", "kubectl")

	kubectlArgs := []string{
		"-n",
		nameSpace,
		"port-forward",
		"svc/" + service,
		"" + strconv.Itoa(targetPort) + ":" + strconv.Itoa(port),
	}

	log.Printf("Executing %s command: %s %s", kubectlCommand, kubectlCommand, strings.Join(kubectlArgs, " "))

	kubectlPortForward := exec.Command(kubectlCommand, kubectlArgs...)

	return kubectlPortForward.Start()
}
