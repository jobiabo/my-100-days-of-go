package main

import (
	"os"
	"fmt"
)
func main() {
	args := os.Args 
	switch len(args) {
	case 2:
		// input := args[1]

		fonts, err := Loadbanner("standard.txt")
		if err != nil {
			fmt.Println(err)
			return 
		}

		validbanner := ValidateBanner(fonts) 
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(validbanner)
	}
}