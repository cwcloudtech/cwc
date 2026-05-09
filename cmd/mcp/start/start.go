package start

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
	port     int
	endpoint string
	listenAddr string
)

type runCwcCommandArgs struct {
	Command string   `json:"command" jsonschema:"required,description=The cwc command to run without the leading cwc binary name"`
	Args    []string `json:"args" jsonschema:"description=Additional command arguments and flags"`
}

// StartCmd runs a stateless MCP HTTP server.
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the cwc MCP server",
	Long:  "Start the cwc MCP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		addr := fmt.Sprintf("%s:%d", listenAddr, port)
		transport := mcp_http_transport.NewHTTPTransport(endpoint).WithAddr(addr)

		server := mcp_golang.NewServer(
			transport,
			mcp_golang.WithName("cwc-mcp-server"),
			mcp_golang.WithVersion("0.1.0"),
			mcp_golang.WithInstructions("Use run_cwc_command to execute cwc CLI commands."),
		)

		err := server.RegisterTool(
			"run_cwc_command",
			"Run any available cwc CLI command. Pass command and args without the leading cwc binary name.",
			func(arguments runCwcCommandArgs) (*mcp_golang.ToolResponse, error) {
				commandName := strings.TrimSpace(arguments.Command)
				if commandName == "" {
					return nil, fmt.Errorf("command is required")
				}

				if commandName == "mcp" && len(arguments.Args) > 0 && arguments.Args[0] == "start" {
					return nil, fmt.Errorf("running mcp start from the MCP tool is blocked")
				}
				if commandName == "ai" && len(arguments.Args) > 1 && arguments.Args[0] == "mcp" && arguments.Args[1] == "start" {
					return nil, fmt.Errorf("running ai mcp start from the MCP tool is blocked")
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
	StartCmd.Flags().StringVarP(&listenAddr, "listen", "l", "127.0.0.1", "MCP server listen address")
	StartCmd.Flags().IntVarP(&port, "port", "p", 8080, "MCP server port")
	StartCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "/mcp", "MCP HTTP endpoint path")
}
