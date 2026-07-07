// EXERCISE 15
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 2 3 4"
	fmt.Println(SumLoops(input))

}

func SumLoops(input string) int {
	total := 0

	// slice of string
	splitS := strings.Fields(input)

	for i := 0; i < len(splitS); i++ {
		// here each slice of string is converted to a int
		num, _ := strconv.Atoi(splitS[i])
		if num%2 == 0 {
			// here we are returning the sum as int
			total += num
		}
	}
	return total
}
