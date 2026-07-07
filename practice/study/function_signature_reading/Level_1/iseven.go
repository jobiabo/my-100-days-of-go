// EXERCISE 4

package main

import "fmt"

func main() {
	fmt.Println(IsEven(4))
	fmt.Println(IsEven(37))

	
}

func IsEven(n int) bool {
	if n % 2 == 0 {
		return true
	}
	return false
}