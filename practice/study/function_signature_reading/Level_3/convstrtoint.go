// EXERCISE 11
package main

import (
	"fmt"
	"strconv"
)

func main() {
	a, _ := strconv.Atoi("42")
	fmt.Println(a)
	b, _ := strconv.Atoi("123")
	fmt.Println(b)
	c := strconv.Itoa(b)
	fmt.Println(c)

}
