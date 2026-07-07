package ascii

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {

	parts := SplitInput(input)
	// fmt.Println(parts)

	var final []string

	for _, part := range parts {

		rendered := RenderBanner(part, banner)

		final = append(final, rendered...)

	}

	return strings.Join(final, "\n")
}
