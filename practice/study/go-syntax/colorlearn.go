package main

import "fmt"

func main() {
	// red := "\033[31m"
	// reset := "\033[0m"

	// fmt.Println(red)
	fmt.Println("hello" + "\033[31m")
	// fmt.Println(reset)
	fmt.Println("hi")

	a := "\033[31m" + "hello"
	b := "\033[32m" + "World"

	slice := []string{}

	slice = append(slice, a)
	slice = append(slice, b)

	fmt.Println(slice)

}
