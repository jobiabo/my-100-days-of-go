package main

import (
	"ascii/ascii"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please include text to transform")
		return
	}

	switch len(os.Args) {
	case 2:
		input := os.Args[1]

		_, err := ascii.ValidateInput(input)
		if err != nil {
			fmt.Println("invalid input")
			return
		}

		draft, err := ascii.LoadBanner("standard.txt")
		if err != nil {
			fmt.Println(err)
			return
		}

		// for i := 32; i <= 126; i++ {
		// 	fmt.Println(strings.Join((draft[rune(i)]), "\n"))
		// }
		g := ascii.GenerateArt(input, draft)
		fmt.Println(g)

	}
}
