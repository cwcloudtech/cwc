package mcp

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	mcp_golang "github.com/metoro-io/mcp-golang"
	mcp_http_transport "github.com/metoro-io/mcp-golang/transport/http"
	"github.com/spf13/cobra"
)

var (
	port       int
	endpoint   string
	listenAddr string
)

type runCwcCommandArgs struct {
	Command string   `json:"command" jsonschema:"required,description=The cwc command to run without the leading cwc binary name"`
	Args    []string `json:"args" jsonschema:"description=Additional command arguments and flags"`
}

type getCwcCommandHelpArgs struct {
	Command string `json:"command" jsonschema:"required,description=Top-level cwc command to get help for (e.g. instance, project, environment)"`
}

// McpCmd represents the MCP command group under ai.
var McpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start the cwc MCP server",
	Long:  "Start the cwc MCP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		addr := fmt.Sprintf("%s:%d", listenAddr, port)
		transport := mcp_http_transport.NewHTTPTransport(endpoint).WithAddr(addr)

		server := mcp_golang.NewServer(
			transport,
			mcp_golang.WithName("cwc-mcp-server"),
			mcp_golang.WithVersion("0.1.0"),
			mcp_golang.WithInstructions("Use list_cwc_commands and get_cwc_command_help to discover valid CLI commands, then execute with run_cwc_command."),
		)

		err := server.RegisterTool(
			"list_cwc_commands",
			"List top-level cwc commands by returning `cwc --help` output. Use this first before calling run_cwc_command.",
			func(arguments struct{}) (*mcp_golang.ToolResponse, error) {
				executable, err := os.Executable()
				if err != nil {
					return nil, fmt.Errorf("failed to resolve cwc executable: %w", err)
				}

				runCmd := exec.Command(executable, "--help")
				output, err := runCmd.CombinedOutput()

				exitCode := 0
				if runCmd.ProcessState != nil {
					exitCode = runCmd.ProcessState.ExitCode()
				}

				result := fmt.Sprintf("command: cwc --help\nexit_code: %d\noutput:\n%s", exitCode, string(output))
				if err != nil {
					return nil, fmt.Errorf("%s", result)
				}

				return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(result)), nil
			},
		)
		if err != nil {
			return err
		}

		err = server.RegisterTool(
			"get_cwc_command_help",
			"Get help for a specific top-level cwc command by returning `cwc <command> --help` output.",
			func(arguments getCwcCommandHelpArgs) (*mcp_golang.ToolResponse, error) {
				commandName := strings.TrimSpace(arguments.Command)
				if commandName == "" {
					return nil, fmt.Errorf("command is required")
				}

				executable, err := os.Executable()
				if err != nil {
					return nil, fmt.Errorf("failed to resolve cwc executable: %w", err)
				}

				runCmd := exec.Command(executable, commandName, "--help")
				output, err := runCmd.CombinedOutput()

				exitCode := 0
				if runCmd.ProcessState != nil {
					exitCode = runCmd.ProcessState.ExitCode()
				}

				result := fmt.Sprintf("command: cwc %s --help\nexit_code: %d\noutput:\n%s", commandName, exitCode, string(output))
				if err != nil {
					return nil, fmt.Errorf("%s", result)
				}

				return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(result)), nil
			},
		)
		if err != nil {
			return err
		}

		err = server.RegisterTool(
			"run_cwc_command",
			"Run a valid cwc CLI command. Pass command and args without the leading cwc binary name. Use list_cwc_commands/get_cwc_command_help first to avoid invalid command names.",
			func(arguments runCwcCommandArgs) (*mcp_golang.ToolResponse, error) {
				commandName := strings.TrimSpace(arguments.Command)
				if commandName == "" {
					return nil, fmt.Errorf("command is required")
				}

				if commandName == "ai" && len(arguments.Args) > 0 && arguments.Args[0] == "mcp" {
					return nil, fmt.Errorf("running ai mcp from the MCP tool is blocked")
				}

				cliArgs := append([]string{commandName}, arguments.Args...)

				executable, err := os.Executable()
				if err != nil {
					return nil, fmt.Errorf("failed to resolve cwc executable: %w", err)
				}

				runCmd := exec.Command(executable, cliArgs...)
				output, err := runCmd.CombinedOutput()

				exitCode := 0
				if runCmd.ProcessState != nil {
					exitCode = runCmd.ProcessState.ExitCode()
				}

				result := fmt.Sprintf("command: cwc %s\nexit_code: %d\noutput:\n%s", strings.Join(cliArgs, " "), exitCode, string(output))
				if err != nil {
					return nil, fmt.Errorf("%s", result)
				}

				return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(result)), nil
			},
		)
		if err != nil {
			return err
		}

		fmt.Printf("Starting cwc MCP server on http://%s%s\n", addr, endpoint)
		return server.Serve()
	},
}

func init() {
	McpCmd.DisableFlagsInUseLine = true
	McpCmd.Flags().StringVarP(&listenAddr, "listen", "l", "127.0.0.1", "MCP server listen address")
	McpCmd.Flags().IntVarP(&port, "port", "p", 8080, "MCP server port")
	McpCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "/mcp", "MCP HTTP endpoint path")
}
