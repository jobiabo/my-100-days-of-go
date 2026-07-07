// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	substring := "kit"
// 	word := "a kitten in the kit"

// 	if strings.Contains(word, substring) {
// 		fmt.Printf("Found substring '%s' in the word!\n", substring)

// 		for i := 0; i < len(word); i++ {
// 			if word[i] == substring[0] && word[i+1] == substring[1] && word[i+2] == substring[2] {
// 				fmt.Println(i, string(word[i]))
// 				fmt.Println(i, string(word[i+1]))
// 				fmt.Println(i, string(word[i+2]))

// 			}
// 		}
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	word := "a kitten in the kit"
// 	substring := "kit"

// 	// Check if substring is inside the word first
// 	if strings.Contains(word, substring) {
// 		fmt.Printf("Found substring '%s' in the word!\n", substring)

// 		// Range loop over each character
// 		for i, char := range word {
// 			if string(char) == "k" {
// 				fmt.Println(i, string(char))
// 			}
// 		}
// 	}
// }

package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "a kitten in the kit"
	sub := "kit"

	fmt.Println("position:", (MatchString(word, sub)))
}

func MatchString(word string, substring string) []int {
	position := []int{}

	offset := 0

	for {
		i := strings.Index(word[offset:], substring)
		if i == -1 {
			break
		}

		start := offset + i
		for j := 0; j < len(substring); j++ {
			position = append(position, start+j)
			fmt.Println(start + j)
		}

		offset = start + len(substring)
	}
	return position
}
