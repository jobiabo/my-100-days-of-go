package main

import (
	"converter/converter"
	"fmt"
)

func main() {
	g := converter.StringToArt("1")
	fmt.Println(g)
	fmt.Println(len(g))
}
