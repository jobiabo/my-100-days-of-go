package main

import "fmt"

func firstUniqChar(s string) int {
	charStore := map[rune]int{}

	for _, char := range s {
		charStore[(char)]++

	}
	for i, ch := range s {
		if charStore[(ch)] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	input := "leetcode"

	fmt.Println(firstUniqChar(input))
}
