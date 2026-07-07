package ascii

import (
	"strings"
)

func RenderBanner(input string, banner map[rune][]string) []string {
	result := []string{}

	for row := 0; row < 8; row++ {
		var draw strings.Builder
		for _, char := range input {
			draw.WriteString(banner[char][row])
		}
		result = append(result, draw.String())
	}
	return result
}