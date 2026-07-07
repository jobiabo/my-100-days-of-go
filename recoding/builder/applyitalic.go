package builder

import "strings"

func ApplyItalic(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	result := []string{}

	for i, line := range lines {
		result = append(result, strings.Repeat(" ", i)+line)
	}
	return result
}
