package ascii 

import "strings"

func GenerateArt(input string, banner map[rune][]string) string {
	output := []string{}
	parts := SplitInput(input)

	for _, part := range parts {
		rendered := RenderBanner(part, banner)

		output = append(output, rendered...)
	}
	return strings.Join(output, "\n")
}