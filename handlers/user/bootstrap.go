package user

import (
	"cwc/env"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"cwc/utils"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/spf13/cobra"
)

type RepoConfig struct {
	RepoURL  string
	Branch   string
	Username string
	Password string
}

func GetRepoConfig() RepoConfig {
	return RepoConfig{
		RepoURL: env.REPO_URL,
		Branch:  env.BRANCH,
	}
}

func HandleTemporaryConfig(tempConfig *RepoConfig) (cleanup func()) {
	if tempConfig == nil {
		return func() {}
	}

	originalConfig := GetRepoConfig()

	if tempConfig.RepoURL != "" {
		env.REPO_URL = tempConfig.RepoURL
	}
	if tempConfig.Branch != "" {
		env.BRANCH = tempConfig.Branch
	}

	return func() {
		env.REPO_URL = originalConfig.RepoURL
		env.BRANCH = originalConfig.Branch
	}
}

func HandleBootstrap(cmd *cobra.Command, releaseName, nameSpace string, otherValues []string, flagVerbose bool, keepDir bool, recreateNs bool, openshift bool) {
	repoURL := env.REPO_URL
	directory := env.DIRECTORY
	branch := env.BRANCH

	if err := CloneRepo(repoURL, directory, branch, keepDir, "", ""); err != nil {
		log.Printf("Error cloning repository: %v", err)
		return
	}

	log.Println("Starting Helm chart installation...")

	patchString := buildPatchString(otherValues)
	log.Printf("Constructed patch string: %s", patchString)

	if err := runHelmDependancyUpdate(directory, keepDir); err != nil {
		log.Fatalf("Error running helm command: %v", err)
	}

	if err := runDeleteNS(nameSpace, recreateNs, openshift); err != nil {
		log.Printf("Not able to delete the namespace: %s, error: %v", nameSpace, err)
	}

	if err := runHelmInstall(releaseName, directory, nameSpace, patchString, openshift); err != nil {
		log.Fatalf("Error running helm command: %v", err)
	}

	log.Println("Helm chart installation completed successfully.")
}

func HandleBootstrapWithConfig(cmd *cobra.Command, releaseName, nameSpace string, otherValues []string, flagVerbose bool, keepDir bool, openshift bool, tempConfig *RepoConfig) {
	cleanup := HandleTemporaryConfig(tempConfig)
	defer cleanup()

	if err := CloneRepo(tempConfig.RepoURL, env.DIRECTORY, tempConfig.Branch, keepDir, tempConfig.Username, tempConfig.Password); err != nil {
		log.Printf("Error cloning repository: %v", err)
		return
	}

	log.Println("Starting Helm chart installation...")

	patchString := buildPatchString(otherValues)
	log.Printf("Constructed patch string: %s", patchString)

	if err := runHelmInstall(releaseName, env.DIRECTORY, nameSpace, patchString, openshift); err != nil {
		log.Fatalf("Error running helm command: %v", err)
	}

	log.Println("Helm chart installation completed successfully.")
}

func buildPatchString(otherValues []string) string {
	clusterIP := GetClusterIP()

	var builder strings.Builder
	builder.WriteString("clusterIP=" + clusterIP)

	for _, opt := range otherValues {
		builder.WriteString("," + opt)
	}

	return builder.String()
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

func runHelmDependancyUpdate(directory string, keepDir bool) error {
	if keepDir {
		return nil
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

func runHelmInstall(releaseName, directory, nameSpace, patchString string, openshift bool) error {
	helmCommand := "helm"
	helmArgs := []string{
		"install",
		releaseName,
		directory,
		utils.If(openshift, "", "--create-namespace"),
		"--namespace", nameSpace,
		utils.If(openshift, "--set", ""),
		utils.If(openshift, "s3.enabled=false", ""),
		"--set",
		patchString,
	}

	log.Printf("Executing helm command: %s %s", helmCommand, strings.Join(helmArgs, " "))

	helmInstallation := exec.Command(helmCommand, helmArgs...)
	helmInstallation.Stdout = os.Stdout
	helmInstallation.Stderr = os.Stderr

	return helmInstallation.Run()
}

func CloneRepo(repoURL, directory, branch string, keepDir bool, username, password string) error {
	if !keepDir {
		if _, err := os.Stat(directory); !os.IsNotExist(err) {
			fmt.Printf("Deleting existing directory: %s\n", directory)
			if err := os.RemoveAll(directory); err != nil {
				return fmt.Errorf("failed to delete existing directory: %v", err)
			}
		}
	}

	cloneOptions := &git.CloneOptions{
		URL:           repoURL,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		Progress:      os.Stdout,
	}

	if username != "" || password != "" {
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

func HandlePortForward(cmd *cobra.Command, nameSpace string, openshift bool) {
	log.Println("Starting tunnel on CWCloud...")

	if err := runPortForward(nameSpace, "api", 8000, openshift); err != nil {
		log.Fatalf("Error running kubectl: %v", err)
	}

	if err := runPortForward(nameSpace, "ui", 3000, openshift); err != nil {
		log.Fatalf("Error running kubectl: %v", err)
	}

	log.Println("Now you can go here: http://localhost:3000")
}

func runPortForward(nameSpace string, service string, port int, openshift bool) error {
	kubectlCommand := utils.If(openshift, "oc", "kubectl")

	kubectlArgs := []string{
		"-n",
		nameSpace,
		"port-forward",
		"svc/cwcloud-" + service,
		"" + strconv.Itoa(port) + ":" + strconv.Itoa(port),
	}

	log.Printf("Executing %s command: %s %s", kubectlCommand, kubectlCommand, strings.Join(kubectlArgs, " "))

	kubectlPortForward := exec.Command(kubectlCommand, kubectlArgs...)

	return kubectlPortForward.Start()
}
