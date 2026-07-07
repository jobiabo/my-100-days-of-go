package ascii

import (
	"strings"
)


func Generate(input string, banner map[rune][]string) string {
	result := []string{}
	parts := SplitInput(input)

	for _, part := range parts {
		rendered := Render(part, banner)
		result = append(result, rendered...)
	}
	return strings.Join(result, "\n")
} 