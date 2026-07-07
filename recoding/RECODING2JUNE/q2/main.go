package main

import (
	"fmt"
	"converter/converter"
)

func main() {
	a := "1\n2"
	fmt.Println(converter.StringToArt(a))
}