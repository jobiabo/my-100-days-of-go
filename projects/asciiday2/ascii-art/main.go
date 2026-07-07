package main

import (
	"fmt"
	"generator/generator"
)

func main() {
	a := 'Z'
	fmt.Println(generator.GeneratePattern(a))
}