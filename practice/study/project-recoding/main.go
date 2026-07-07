package main

import (
	"fmt"
)

func main() {
	input := "" +
		" ___ \n" +
		"|   |\n" +
		"|   |\n" +
		"|   |\n" +
		"|___|"
	// d := strings.Join(input, "\n")
	fmt.Println(ArtToString(input))
}
