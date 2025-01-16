package bootstrap

import (
	"cwc/cmd/bootstrap/pfw"
	"cwc/cmd/bootstrap/uninstall"
	"cwc/env"
	"cwc/handlers/user"
	"cwc/utils"
	"fmt"
	"github.com/spf13/cobra"
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
	BootstrapCmd.Flags().StringArrayVarP(&otherValues, "value", "p", []string{}, `Values to override other configurations (e.g. --value key=value --value key2=value2)`)

	configureCmd := &cobra.Command{
		Use:   "configure",
		Short: "Bootstrap CWCloud with custom repository configuration",
		Long: `Bootstrap CWCloud installation with temporary repository configuration overrides. 
		This command allows you to specify custom repository URL, directory, and branch for a single deployment.

		You can either:
		1. Use flags to specify all values directly
		2. Run without flags for an interactive prompt`,
		Example: `  # Interactive mode:
		cwc bootstrap configure

		# Flag-based mode:
		cwc bootstrap configure --repo-url=https://custom-repo.git --branch=dev
		cwc bootstrap configure --repo-url=https://gitlab.alternative.io/helm.git --branch=feature-123
		cwc bootstrap configure --repo-url=https://custom-repo.git --username=myuser --password=mypassword`,
		Run: func(cmd *cobra.Command, args []string) {
			defaultRepoURL := env.REPO_URL
			defaultBranch := env.BRANCH

			if !cmd.Flags().Changed("repo-url") && !cmd.Flags().Changed("branch") &&
				!cmd.Flags().Changed("username") && !cmd.Flags().Changed("password") {

				fmt.Printf("Repository URL [%s]: ", defaultRepoURL)
				userRepoURL := utils.PromptUserForValue()
				if utils.IsNotBlank(userRepoURL) {
					tempRepoURL = userRepoURL
				} else {
					tempRepoURL = defaultRepoURL
				}

				fmt.Printf("Branch [%s]: ", defaultBranch)
				userBranch := utils.PromptUserForValue()
				if utils.IsNotBlank(userBranch) {
					tempBranch = userBranch
				} else {
					tempBranch = defaultBranch
				}

				if tempRepoURL != defaultRepoURL {
					fmt.Print("Username (optional): ")
					tempUsername = utils.PromptUserForValue()

					if utils.IsNotBlank(tempUsername) {
						fmt.Print("Password (optional): ")
						tempPassword = utils.PromptUserForValue()
					}
				}

			} else {
				if !cmd.Flags().Changed("repo-url") {
					tempRepoURL = defaultRepoURL
				}
				if !cmd.Flags().Changed("branch") {
					tempBranch = defaultBranch
				}
			}

			tempConfig := &user.RepoConfig{
				RepoURL:  tempRepoURL,
				Branch:   tempBranch,
				Username: tempUsername,
				Password: tempPassword,
			}
			user.HandleBootstrapWithConfig(cmd, releaseName, nameSpace, otherValues, flagVerbose, keepDir, openshift, tempConfig)
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
