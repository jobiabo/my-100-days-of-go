// EXERCISE 1

package main

import (
	"fmt"
	"strings"

)

func main() {
	a := "aeiouplk" // expecting 5 here
	b := "hello to the world" // expecting 2 here
	fmt.Println(VowelsCount(a))
	fmt.Println(VowelsCount(b))

	
}

func VowelsCount(text string) int {
	vowels := "aeiou"
	seen := make(map[rune]bool)

	text = strings.ToLower(text)

	for _, ch := range text {
		if strings.ContainsRune(vowels, ch) {
			seen[ch] = true
		}
	}

	return len(seen)
	
}
