package agent

import (
	mcpPrompt "cwc/cmd/mcp/prompt"
)

// AgentCmd delegates to the MCP-backed AI agent prompt command.
var AgentCmd = mcpPrompt.PromptCmd

func init() {
	AgentCmd.Use = "agent"
	AgentCmd.Short = "Send a prompt to an AI agent backed by MCP server tools"
	AgentCmd.Long = "Send a prompt to an AI agent that uses MCP server tools to execute the request"
}
