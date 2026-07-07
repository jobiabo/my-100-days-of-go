package ascii

import "strings"

// func RenderBanner(input string, banner map[rune][]string) []string {
// 	color := Color(cmd.Color)
// 	// position := MatchString(cmd.Text, cmd.Substring)

// 	result := []string{}

// 	for row := 0; row < 8; row++ {
// 		var draw strings.Builder
// 		for i, char := range input {
// 			if i == 8 {
// 				g := banner[char]
// 				art := strings.Join(g, "\n") + color
// 				draw.WriteString(art)
// 			}
// 		}
// 		result = append(result, draw.String())

// 	}
// 	return result
// }

// func RenderBanner(input string, substring string, banner map[rune][]string) []string {
// 	sen := strings.Fields(input)
// 	result := []string{}
// 	lenOfSubStr := 0
// 	for _, word := range sen {
// 		for row := 0; row < 8; row++ {
// 			var draw strings.Builder
// 			for i, char := range input {

// 				if strings.HasPrefix(word[i:], substring) {
// 					lenOfSubStr = len(substring)
// 				}

// 				if lenOfSubStr > 0 {

// 					lenOfSubStr--
// 				} else {
// 					draw.WriteString(banner[char][row])
// 				}
// 			}
// 			result = append(result, draw.String())

// 		}
// 	}
// 	return result
// }

func RenderBanner(input string, banner map[rune][]string) []string {
	color := Color(cmd.Color)
	reset := "\033[0m"

	positions := MatchString(input, cmd.Substring)

	highlight := make(map[int]bool)

	for _, start := range positions {
		for i := 0; i < len(cmd.Substring); i++ {
			highlight[start+i] = true
		}
	}

	result := make([]string, 0, 8)

	for row := 0; row < 8; row++ {
		var draw strings.Builder

		for i, char := range input {
			art := banner[char][row]

			if highlight[i] {
				draw.WriteString(color)
				draw.WriteString(art)
				draw.WriteString(reset)
			} else {
				draw.WriteString(art)
			}
		}

		result = append(result, draw.String())
	}

	return result
}
