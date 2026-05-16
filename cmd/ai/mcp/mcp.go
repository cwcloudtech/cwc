package mcp

import (
	"cwc/utils"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
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

type dynamicCommandArgs struct {
	Args []string `json:"args" jsonschema:"description=Additional args, subcommands and flags for this command path"`
}

type emptyToolArgs struct{}

type commandSpec struct {
	Path        []string
	Description string
}

// McpCmd represents the MCP command group under ai.
var McpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Start the cwc MCP server",
	Long:  "Start the cwc MCP server",
	RunE: func(cmd *cobra.Command, args []string) error {
		addr := fmt.Sprintf("%s:%d", listenAddr, port)
		transport := mcp_http_transport.NewHTTPTransport(endpoint).WithAddr(addr)

		executable, err := os.Executable()
		if err != nil {
			return fmt.Errorf("failed to resolve cwc executable: %w", err)
		}

		dynamicCommands, err := discoverCWCCommands(executable)
		if err != nil {
			return fmt.Errorf("failed to discover cwc commands: %w", err)
		}

		dynamicToolLines := make([]string, 0, len(dynamicCommands))
		for _, spec := range dynamicCommands {
			if len(spec.Path) == 0 {
				continue
			}
			if len(spec.Path) >= 2 && spec.Path[0] == "ai" && spec.Path[1] == "mcp" {
				continue
			}
			toolName := "cwc_" + sanitizeToolName(strings.Join(spec.Path, "_"))
			desc := strings.TrimSpace(spec.Description)
			if utils.IsBlank(desc) {
				desc = "(no description)"
			}

			dynamicToolLines = append(dynamicToolLines, fmt.Sprintf("- %s => cwc %s | %s", toolName, strings.Join(spec.Path, " "), desc))
		}
		sort.Strings(dynamicToolLines)

		server := mcp_golang.NewServer(
			transport,
			mcp_golang.WithName("cwc-mcp-server"),
			mcp_golang.WithVersion("0.1.0"),
			mcp_golang.WithInstructions("Use list_cwc_commands/get_cwc_command_help to discover commands. Prefer dynamic tools named cwc_<command_path> for direct command execution."),
		)

		err = server.RegisterTool(
			"list_cwc_commands",
			"List top-level cwc commands by returning `cwc --help` output. Use this first before calling run_cwc_command.",
			func(arguments emptyToolArgs) (*mcp_golang.ToolResponse, error) {
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
				if utils.IsBlank(commandName) {
					return nil, fmt.Errorf("command is required")
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
			"list_mcp_dynamic_tools",
			"List all dynamically generated MCP tools and their mapped cwc command paths.",
			func(arguments emptyToolArgs) (*mcp_golang.ToolResponse, error) {
				if len(dynamicToolLines) == 0 {
					return mcp_golang.NewToolResponse(mcp_golang.NewTextContent("No dynamic tools discovered.")), nil
				}
				result := fmt.Sprintf("discovered_dynamic_tools: %d\n%s", len(dynamicToolLines), strings.Join(dynamicToolLines, "\n"))
				return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(result)), nil
			},
		)
		if err != nil {
			return err
		}

		for _, spec := range dynamicCommands {
			if len(spec.Path) == 0 {
				continue
			}

			if len(spec.Path) >= 2 && spec.Path[0] == "ai" && spec.Path[1] == "mcp" {
				continue
			}

			toolName := "cwc_" + sanitizeToolName(strings.Join(spec.Path, "_"))
			desc := strings.TrimSpace(spec.Description)
			if utils.IsBlank(desc) {
				desc = fmt.Sprintf("Run command path: cwc %s", strings.Join(spec.Path, " "))
			}

			fullPath := append([]string{}, spec.Path...)

			err = server.RegisterTool(
				toolName,
				fmt.Sprintf("%s. Base command: cwc %s", desc, strings.Join(fullPath, " ")),
				func(arguments dynamicCommandArgs) (*mcp_golang.ToolResponse, error) {
					cliArgs := append(append([]string{}, fullPath...), arguments.Args...)
					result, runErr := runCLI(executable, cliArgs)
					if runErr != nil {
						return nil, fmt.Errorf("%s", result)
					}
					return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(result)), nil
				},
			)
			if err != nil {
				return err
			}
		}

		err = server.RegisterTool(
			"run_cwc_command",
			"Fallback generic command runner. Prefer dynamic cwc_<command_path> tools discovered from Cobra tree.",
			func(arguments runCwcCommandArgs) (*mcp_golang.ToolResponse, error) {
				commandName := strings.TrimSpace(arguments.Command)
				cliArgs := make([]string, 0, 1+len(arguments.Args))
				if utils.IsNotBlank(commandName) {
					cliArgs = append(cliArgs, strings.Fields(commandName)...)
				}

				cliArgs = append(cliArgs, arguments.Args...)

				for len(cliArgs) > 0 && strings.EqualFold(strings.TrimSpace(cliArgs[0]), "cwc") {
					cliArgs = cliArgs[1:]
				}

				if len(cliArgs) == 0 {
					return nil, fmt.Errorf("command is required")
				}

				commandName = strings.TrimSpace(cliArgs[0])
				commandArgs := make([]string, 0, len(cliArgs)-1)
				if len(cliArgs) > 1 {
					commandArgs = append(commandArgs, cliArgs[1:]...)
				}

				if len(commandArgs) > 0 {
					verb := strings.ToLower(strings.TrimSpace(commandArgs[0]))
					switch verb {
					case "ls", "list", "liste", "lister", "show", "display", "give", "montrer":
						commandArgs[0] = "ls"
					case "reboot", "restart", "redemarre":
						commandArgs[0] = "restart"
					case "delete", "remove", "supprimer", "supprime", "efface", "effacer", "rm":
						commandArgs[0] = "delete"
					case "update", "modifier", "modifie":
						commandArgs[0] = "update"
					case "create", "creer", "cree", "ajoute", "ajouter", "new":
						commandArgs[0] = "create"
					}
				}

				if strings.HasSuffix(strings.ToLower(commandName), "s") && len(commandName) > 1 {
					commandName = strings.TrimSuffix(commandName, "s")
				}

				if commandName == "ai" && len(arguments.Args) > 0 && arguments.Args[0] == "mcp" {
					return nil, fmt.Errorf("running ai mcp from the MCP tool is blocked")
				}

				cliArgs = append([]string{commandName}, commandArgs...)

				result, err := runCLI(executable, cliArgs)
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

func runCLI(executable string, cliArgs []string) (string, error) {
	runCmd := exec.Command(executable, cliArgs...)
	output, err := runCmd.CombinedOutput()

	exitCode := 0
	if runCmd.ProcessState != nil {
		exitCode = runCmd.ProcessState.ExitCode()
	}

	result := fmt.Sprintf("command: cwc %s\nexit_code: %d\noutput:\n%s", strings.Join(cliArgs, " "), exitCode, string(output))
	if err != nil {
		return result, err
	}
	return result, nil
}

func discoverCWCCommands(executable string) ([]commandSpec, error) {
	type queueEntry struct {
		Path []string
	}

	queue := []queueEntry{{Path: []string{}}}
	visited := map[string]bool{"": true}
	collected := map[string]commandSpec{}

	for len(queue) > 0 {
		entry := queue[0]
		queue = queue[1:]

		helpArgs := append([]string{}, entry.Path...)
		helpArgs = append(helpArgs, "--help")
		runCmd := exec.Command(executable, helpArgs...)
		output, err := runCmd.CombinedOutput()
		if err != nil {
			continue
		}

		subCommands := parseAvailableCommands(string(output))
		for _, sub := range subCommands {
			path := append(append([]string{}, entry.Path...), sub.Name)
			key := strings.Join(path, " ")
			if !visited[key] {
				visited[key] = true
				queue = append(queue, queueEntry{Path: path})
			}
			collected[key] = commandSpec{Path: path, Description: sub.Description}
		}
	}

	result := make([]commandSpec, 0, len(collected))
	keys := make([]string, 0, len(collected))
	for key := range collected {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		result = append(result, collected[key])
	}

	return result, nil
}

type parsedCommand struct {
	Name        string
	Description string
}

func parseAvailableCommands(helpText string) []parsedCommand {
	lines := strings.Split(helpText, "\n")
	commands := make([]parsedCommand, 0)
	inAvailable := false
	re := regexp.MustCompile(`^\s{2,}([a-zA-Z0-9_-]+)\s{2,}(.*)$`)

	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if !inAvailable {
			if strings.HasPrefix(trim, "Available Commands:") {
				inAvailable = true
			}
			continue
		}

		if utils.IsBlank(trim) {
			continue
		}

		if strings.HasSuffix(trim, ":") {
			break
		}

		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			continue
		}

		name := strings.TrimSpace(matches[1])
		if name == "help" {
			continue
		}

		commands = append(commands, parsedCommand{Name: name, Description: strings.TrimSpace(matches[2])})
	}

	return commands
}

func sanitizeToolName(value string) string {
	value = strings.ReplaceAll(strings.ReplaceAll(value, "-", "_"), " ", "_")

	b := strings.Builder{}
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_' {
			b.WriteRune(r)
		}
	}
	out := b.String()
	if utils.IsBlank(out) {
		return "command"
	}

	return out
}
