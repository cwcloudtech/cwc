package web_agent

import (
	"cwc/utils"
	"strings"
	"testing"
)

func TestFormatWebAgentMessageReturnsRawWhenNotCommandOutput(t *testing.T) {
	raw := "plain response"
	got := utils.FormatWebAgentMessage(raw)

	if got != raw {
		t.Fatalf("expected raw response to be unchanged, got %q", got)
	}
}

func TestFormatWebAgentMessageFormatsConsoleOutput(t *testing.T) {
	raw := strings.Join([]string{
		"command: cwc project list",
		"exit_code: 0",
		"output:",
		"project-a",
		"project-b",
	}, "\n")

	got := utils.FormatWebAgentMessage(raw)

	if !strings.HasPrefix(got, "```console\n$ cwc project list\n") {
		t.Fatalf("expected console fence with command, got %q", got)
	}
	if !strings.Contains(got, "project-a\nproject-b") {
		t.Fatalf("expected output lines in formatted message, got %q", got)
	}
	if !strings.HasSuffix(got, "```") {
		t.Fatalf("expected message to end with code fence, got %q", got)
	}
}

func TestFormatWebAgentMessageConvertsPrettyTableToMarkdown(t *testing.T) {
	raw := strings.Join([]string{
		"command: cwc project list",
		"exit_code: 0",
		"output:",
		"+------+--------+",
		"| Name | Region |",
		"+------+--------+",
		"| p1   | eu     |",
		"| p2   | us     |",
		"+------+--------+",
	}, "\n")

	got := utils.FormatWebAgentMessage(raw)

	if !strings.Contains(got, "```console\n$ cwc project list\n```") {
		t.Fatalf("expected command to stay in console fence, got %q", got)
	}
	if strings.Contains(got, "+------+--------+") {
		t.Fatalf("expected pretty table borders to be converted, got %q", got)
	}
	if !strings.Contains(got, "| Name | Region |") {
		t.Fatalf("expected markdown table header, got %q", got)
	}
	if !strings.Contains(got, "| --- | --- |") {
		t.Fatalf("expected markdown table divider, got %q", got)
	}
	if !strings.Contains(got, "| p1 | eu |") || !strings.Contains(got, "| p2 | us |") {
		t.Fatalf("expected markdown table rows, got %q", got)
	}
}
