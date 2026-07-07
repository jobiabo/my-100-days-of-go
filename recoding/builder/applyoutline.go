package builder

import (
	"strings"
)

func ApplyOutline(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	width := len(lines[0])
	boarder := "+" + strings.Repeat("+", width) + "+"

	result := []string{}
	result = append(result, boarder)

	for _, line := range lines {
		result = append(result, "|"+line+"|")
	}
	result = append(result, boarder)

	if len(result) > 8 {
		return result[:8]
	}

	if len(result) < 8 {
		result = append(result, boarder)
	}
	return result
}
