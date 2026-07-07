package ascii

import "strings"

// func Render(input string, banner map[rune][]string) []string {
// 	result := make([]string, 8)

// 	for _, char := range input {
// 		for row := 0; row < 8; row++ {
// 			var draw strings.Builder

// 			// 	for _, char := range input {

// 			draw.WriteString(banner[char][row])

// 			result = append(result, draw.String())
// 		}
// 		// result = append(result, banner[ch][row])
// 	}

// 	// fmt.Println(result)

// 	return result
// }

// func main() {
// 	st := "hello"
// 	fmt.Println(Render(st, LoadBanner()))
// }

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
