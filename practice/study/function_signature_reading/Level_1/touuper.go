// EXERCISE 2

package main

import (
	"fmt"
	"strings"
)

func main() {
	cap := "hello"
	fmt.Println(ToUpper(cap))
	
}

func ToUpper(s string) string {
	a := strings.ToUpper(s)
	return a
}