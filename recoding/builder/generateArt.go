package builder

import "strings"

func GenerateArt(c rune) []string {
	result := []string{}
	for row := 0; row < 8; row++ {
		result = append(result, strings.Repeat(string(c), 4))
	}
	return result
}
