// EXERCISE 3
package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "how many words are here, i guess it is 10"
	fmt.Println(CountWords(sentence))
	
}

func CountWords(s string) int {
	split_text := strings.Fields(s)
	count := 0

	for i := 0; i < len(split_text); i++ {
		count++
	}
	return count
}