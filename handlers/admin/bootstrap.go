package admin

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	// "path/filepath"
	"strings"
	// "helm.sh/helm/v3/pkg/action"
	// "helm.sh/helm/v3/pkg/chart/loader"
	// "helm.sh/helm/v3/pkg/cli"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
	// "k8s.io/cli-runtime/pkg/genericclioptions"
	// "k8s.io/client-go/tools/clientcmd"
	// "k8s.io/client-go/util/homedir"

)
func HandleBootstrap(cmd *cobra.Command, clusterIP, releaseName, dbPassword, nameSpace string, otherValues []string, flagVerbose bool) {
    repoURL := "https://gitlab.comwork.io/oss/cwcloud/cwcloud-helm.git"
    directory := "./cwcloud"
    branch := "feat/add-helm-templates"

    // Clone the helm repository
    if err := CloneRepo(repoURL, directory, branch); err != nil {
        log.Printf("Error cloning repository: %v", err)
        return
    }

    log.Println("Starting Helm chart installation...")

    patchString := buildPatchString(dbPassword, clusterIP, otherValues)
    log.Printf("Constructed patch string: %s", patchString)

    if err := runHelmInstall(releaseName, directory, nameSpace, patchString); err != nil {
        log.Fatalf("Error running helm command: %v", err)
    }

    log.Println("Helm chart installation completed successfully.")
}

func buildPatchString(dbPassword, clusterIP string, otherValues []string) string {
    var builder strings.Builder
    builder.WriteString(fmt.Sprintf("apiEnvCm.postgresPassword=%s,clusterIP=%s", dbPassword, clusterIP))

    for _, opt := range otherValues {
        builder.WriteString("," + opt)
    }

    return builder.String()
}

func runHelmInstall(releaseName, directory, nameSpace, patchString string) error {
    helmCommand := "helm"
    helmArgs := []string{
        "install",
        releaseName,
        directory,
        "--create-namespace",
        "--namespace", nameSpace,
        "--set",
        patchString,
    }

    log.Printf("Executing helm command: %s %s", helmCommand, strings.Join(helmArgs, " "))

    helmInstallation := exec.Command(helmCommand, helmArgs...)
    helmInstallation.Stdout = os.Stdout
    helmInstallation.Stderr = os.Stderr

    return helmInstallation.Run()
}

func CloneRepo(repoURL, directory, branch string) error {

	if _, err := os.Stat(directory); !os.IsNotExist(err) {
        fmt.Printf("Deleting existing directory: %s\n", directory)
        if err := os.RemoveAll(directory); err != nil {
            return fmt.Errorf("failed to delete existing directory: %v", err)
        }
    }

    _, err := git.PlainClone(directory, false, &git.CloneOptions{
        URL:           repoURL,
        ReferenceName: plumbing.NewBranchReferenceName(branch),
        Progress:      os.Stdout,
    })
    if err != nil {
        return fmt.Errorf("failed to clone repository: %v", err)
    }

    fmt.Println("Repository cloned successfully.")
    return nil
}



func HandleUninstall(cmd *cobra.Command, releaseName string, nameSpace string) {
    log.Println("Starting Helm chart uninstallation...")

    // Run helm uninstall command
    if err := runHelmUninstall(releaseName, nameSpace); err != nil {
        log.Fatalf("Error running helm uninstall command: %v", err)
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





// func HandleBootstrap1(cmd *cobra.Command, clusterIP string, releaseName string, dbPassword string, nameSpace string, otherValues []string, flagVerbose bool) {

// 	// Get the namespace of current context and check if it is the same as the namespace provided
// 	currentContextNameSpace, _ := getActiveNamespace()
// 	if currentContextNameSpace != nameSpace {
// 		fmt.Println("\nNamespace in current context is different from the namespace provided.")
// 		fmt.Printf("\nCurrent namespace: %s\n", currentContextNameSpace)
// 		fmt.Printf("Provided namespace: %s\n\n", nameSpace)

// 		fmt.Println("Instructions:")
// 		// We can skip this step but for ensuring that the namespace is created and switched
// 		fmt.Println("1) To create a new namespace, run the following command:")
// 		fmt.Printf("   kubectl create namespace %s\n", nameSpace)

// 		fmt.Println("2) You need to switch to the provided namespace to install the chart application.")
// 		fmt.Printf("   Run the following command:\n   kubectl config set-context --current --namespace=%s\n\n", nameSpace)

// 		return
// 	}

// 	// Clone the helm repository
// 	repoURL := "https://gitlab.comwork.io/oss/cwcloud/cwcloud-helm.git"
// 	directory := "./cwcloud"
// 	branch := "feat/add-helm-templates"

// 	err := CloneRepo(repoURL, directory, branch)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	}

// 	// Install the chart
// 	settings := cli.New()
// 	actionConfig := new(action.Configuration)

// 	if err := actionConfig.Init(genericclioptions.NewConfigFlags(false), nameSpace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
// 		log.Fatalf("Failed to initialize helm configuration: %v", err)
// 	}

// 	client := action.NewInstall(actionConfig)
// 	client.ReleaseName = releaseName
// 	client.CreateNamespace = true
// 	client.Namespace = nameSpace

// 	chartPath, err := client.ChartPathOptions.LocateChart(directory, settings)
// 	if err != nil {
// 		log.Fatalf("Failed to locate chart: %v", err)
// 	}

// 	chart, err := loader.Load(chartPath)
// 	if err != nil {
// 		log.Fatalf("Failed to load chart: %v", err)
// 	}

// 	vals := make(map[string]interface{})

// 	for _, opt := range otherValues {
// 		parts := strings.SplitN(opt, "=", 2)
// 		if len(parts) != 2 {
// 			continue
// 		}

// 		key := strings.TrimPrefix(parts[0], "-p ")
// 		value := parts[1]

// 		// Split the key into parts to create nested maps
// 		keys := strings.Split(key, ".")
// 		current := vals
// 		for i := 0; i < len(keys)-1; i++ {
// 			if _, ok := current[keys[i]]; !ok {
// 				current[keys[i]] = make(map[string]interface{})
// 			}
// 			current = current[keys[i]].(map[string]interface{})
// 		}
// 		current[keys[len(keys)-1]] = value
// 	}
// 	vals["apiEnvCm"] = map[string]interface{}{
// 		"postgresPassword": dbPassword,
// 	}
// 	vals["clusterIP"] = clusterIP
// 	printMap(vals, 0)

// 	release, err := client.Run(chart, vals)
// 	if err != nil {
// 		log.Fatalf("Failed to install chart: %v", err)
// 	}

// 	fmt.Printf("Installed %s chart with status: %s\n", release.Name, release.Info.Status)



// }
// func printMap(m map[string]interface{}, indentLevel int) {
// 	indent := strings.Repeat("\t", indentLevel)
// 	for k, v := range m {
// 		switch val := v.(type) {
// 		case map[string]interface{}:
// 			fmt.Printf("%s%s:\n", indent, k)
// 			printMap(val, indentLevel+1)
// 		default:
// 			fmt.Printf("%s%s: %v\n", indent, k, v)
// 		}
// 	}
// }

// func getActiveNamespace() (string, error) {
// 	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

// 	if envKubeConfig := os.Getenv("KUBECONFIG"); envKubeConfig != "" {
// 		kubeconfigPath = envKubeConfig
// 	}

// 	config, err := clientcmd.LoadFromFile(kubeconfigPath)
// 	if err != nil {
// 		return "", fmt.Errorf("failed to load kubeconfig from %s: %v", kubeconfigPath, err)
// 	}

// 	currentContext := config.CurrentContext
// 	context := config.Contexts[currentContext]
// 	namespace := context.Namespace
// 	if namespace == "" {
// 		namespace = "default"
// 	}

// 	return namespace, nil
// }


