package utils

import "strings"

func FormatWebAgentMessage(raw string) string {
	if !strings.HasPrefix(raw, "command: cwc ") || !strings.Contains(raw, "\noutput:\n") {
		return raw
	}

	command := ""
	exitCode := ""
	lines := strings.Split(raw, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "command: ") {
			command = strings.TrimSpace(strings.TrimPrefix(line, "command: "))
			continue
		}

		if strings.HasPrefix(line, "exit_code: ") {
			exitCode = strings.TrimSpace(strings.TrimPrefix(line, "exit_code: "))
		}
	}

	output := raw
	if idx := strings.Index(raw, "\noutput:\n"); idx >= 0 {
		output = raw[idx+len("\noutput:\n"):]
	}
	output = strings.TrimRight(output, "\n")
	convertedOutput, hasPrettyTable := convertPrettyTablesToMarkdown(output)

	var builder strings.Builder
	if IsNotBlank(command) {
		builder.WriteString("```console\n")
		builder.WriteString("$ ")
		builder.WriteString(command)
		builder.WriteString("\n```")
	}

	if hasPrettyTable && IsNotBlank(convertedOutput) {
		if builder.Len() > 0 {
			builder.WriteString("\n\n")
		}
		builder.WriteString(convertedOutput)
	} else if IsNotBlank(output) {
		builder.WriteString("```console\n")
		builder.WriteString(output)
		builder.WriteString("\n")
		builder.WriteString("```")
	}

	if IsNotBlank(exitCode) && exitCode != "0" {
		if builder.Len() > 0 {
			builder.WriteString("\n\n")
		}
		builder.WriteString("# exit_code: ")
		builder.WriteString(exitCode)
		builder.WriteString("\n")
	}

	return builder.String()
}

func convertPrettyTablesToMarkdown(output string) (string, bool) {
	lines := strings.Split(output, "\n")
	convertedLines := make([]string, 0, len(lines))
	hasPrettyTable := false

	for i := 0; i < len(lines); {
		tableMarkdown, consumed, ok := parsePrettyTable(lines, i)
		if ok {
			convertedLines = append(convertedLines, tableMarkdown)
			hasPrettyTable = true
			i += consumed
			continue
		}

		convertedLines = append(convertedLines, lines[i])
		i++
	}

	return strings.Join(convertedLines, "\n"), hasPrettyTable
}

func parsePrettyTable(lines []string, start int) (string, int, bool) {
	if start+2 >= len(lines) || !isPrettyTableBorderLine(lines[start]) || !isPrettyTableRowLine(lines[start+1]) || !isPrettyTableBorderLine(lines[start+2]) {
		return "", 0, false
	}

	headerCells := parsePrettyTableRowCells(lines[start+1])
	if len(headerCells) == 0 {
		return "", 0, false
	}

	dataRows := make([][]string, 0)
	i := start + 3
	for i < len(lines) {
		if isPrettyTableRowLine(lines[i]) {
			row := parsePrettyTableRowCells(lines[i])
			if len(row) != len(headerCells) {
				return "", 0, false
			}
			dataRows = append(dataRows, row)
			i++
			continue
		}

		if isPrettyTableBorderLine(lines[i]) {
			i++
			break
		}

		break
	}

	if len(dataRows) == 0 {
		return "", 0, false
	}

	var builder strings.Builder
	builder.WriteString("| ")
	builder.WriteString(strings.Join(headerCells, " | "))
	builder.WriteString(" |\n")
	builder.WriteString("|")
	for range headerCells {
		builder.WriteString(" --- |")
	}
	builder.WriteString("\n")

	for idx, row := range dataRows {
		if idx > 0 {
			builder.WriteString("\n")
		}
		builder.WriteString("| ")
		builder.WriteString(strings.Join(row, " | "))
		builder.WriteString(" |")
	}

	return builder.String(), i - start, true
}

func isPrettyTableBorderLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) < 3 || trimmed[0] != '+' || trimmed[len(trimmed)-1] != '+' {
		return false
	}

	for _, ch := range trimmed {
		if ch != '+' && ch != '-' {
			return false
		}
	}

	return strings.Contains(trimmed, "-")
}

func isPrettyTableRowLine(line string) bool {
	trimmed := strings.TrimSpace(line)
	return len(trimmed) >= 3 && trimmed[0] == '|' && trimmed[len(trimmed)-1] == '|'
}

func parsePrettyTableRowCells(line string) []string {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) < 2 || trimmed[0] != '|' || trimmed[len(trimmed)-1] != '|' {
		return nil
	}

	parts := strings.Split(trimmed, "|")
	if len(parts) < 3 {
		return nil
	}

	cells := make([]string, 0, len(parts)-2)
	for _, part := range parts[1 : len(parts)-1] {
		cell := strings.TrimSpace(part)
		cell = strings.ReplaceAll(cell, "|", "\\|")
		cells = append(cells, cell)
	}

	return cells
}
