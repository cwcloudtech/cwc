package bootstrap

import (
	"github.com/spf13/cobra"

	"cwc/cmd/bootstrap/pfw"
	"cwc/cmd/bootstrap/uninstall"
	"cwc/handlers/user"
)

var (
	flagVerbose  bool
	nameSpace    string
	otherValues  []string
	releaseName  string
	keepDir      bool
	recreateNs   bool
	openshift    bool
	tempRepoURL  string
	tempBranch   string
	tempUsername string
	tempPassword string
)

var BootstrapCmd = &cobra.Command{
	Use:   "bootstrap [flags]",
	Short: "CWCloud installation on Kubernetes",
	Long:  `CWCloud installation on Kubernetes.`,
	Run: func(cmd *cobra.Command, args []string) {
		user.HandleBootstrap(cmd, releaseName, nameSpace, otherValues, flagVerbose, keepDir, recreateNs, openshift)
	},
}

func init() {
	BootstrapCmd.DisableFlagsInUseLine = true
	BootstrapCmd.Flags().StringVarP(&releaseName, "release", "r", "release-0.1.0", "Release name for deployment (default: release-0.1.0)")
	BootstrapCmd.Flags().StringVarP(&nameSpace, "namespace", "n", "cwcloud", "Namespace to use for deployment (default: cwcloud)")
	BootstrapCmd.Flags().BoolVarP(&keepDir, "keep-dir", "k", false, "Keep the local helm directory")
	BootstrapCmd.Flags().BoolVarP(&recreateNs, "recreate-ns", "d", false, "Recreate the namespace")
	BootstrapCmd.Flags().BoolVarP(&openshift, "openshift", "o", false, "Use openshift cli instead of kubectl")
	BootstrapCmd.Flags().StringArrayVarP(&otherValues, "value", "p", []string{}, `Values to override other configurations (e.g. --value key=value --value key2=value2)

Example:
  All applications are enabled by default. To disable some applications, use this format:

    -p applicationName.enabled=false


  Example:
    cwc bootstrap -p scheduler.enabled=false \
                  -p consumer.enabled=false
	`)

	configureCmd := &cobra.Command{
		Use:   "configure [flags]",
		Short: "Bootstrap CWCloud with custom repository configuration",
		Long: `Bootstrap CWCloud installation with temporary repository configuration overrides.
This command allows you to specify custom repository URL, directory, and branch for a single deployment.`,
		Example: `  cwc bootstrap configure --repo-url=https://custom-repo.git --branch=dev -r my-release -n my-namespace
  cwc bootstrap configure --repo-url=https://gitlab.alternative.io/helm.git --directory=./custom-helm --branch=feature-123
  cwc bootstrap configure --repo-url=https://custom-repo.git --username=myuser --password=mypassword`,
		Run: func(cmd *cobra.Command, args []string) {
			tempConfig := &user.RepoConfig{
				RepoURL:  tempRepoURL,
				Branch:   tempBranch,
				Username: tempUsername,
				Password: tempPassword,
			}
			user.HandleBootstrapWithConfig(cmd, releaseName, nameSpace, otherValues, flagVerbose, keepDir, tempConfig)
		},
	}

	configureCmd.Flags().StringVarP(&releaseName, "release", "r", "release-0.1.0", "Release name for deployment (default: release-0.1.0)")
	configureCmd.Flags().StringVarP(&nameSpace, "namespace", "n", "cwcloud", "Namespace to use for deployment (default: cwcloud)")
	configureCmd.Flags().BoolVarP(&keepDir, "keep-dir", "k", false, "Keep the local helm directory")
	configureCmd.Flags().StringArrayVarP(&otherValues, "value", "p", []string{}, `Values to override other configurations (e.g. --value key=value --value key2=value2)`)

	// Configure subcommand flags
	configureCmd.Flags().StringVarP(&tempRepoURL, "repo-url", "u", "", "Temporary repository URL")
	configureCmd.Flags().StringVarP(&tempBranch, "branch", "b", "", "Temporary branch name")
	configureCmd.Flags().StringVar(&tempUsername, "username", "U", "Username for repository authentication")
	configureCmd.Flags().StringVar(&tempPassword, "password", "P", "Password for repository authentication")

	BootstrapCmd.AddCommand(configureCmd)
	BootstrapCmd.AddCommand(uninstall.UninstallCmd)
	BootstrapCmd.AddCommand(pfw.PfwCmd)
}
