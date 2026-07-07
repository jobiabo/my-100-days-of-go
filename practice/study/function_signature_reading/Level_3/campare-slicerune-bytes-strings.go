// EXERCISE 14
package main

import "fmt"

func main() {
	a := []byte("hello")
	b := []rune("hello")
	c := string([]byte{104, 101, 108, 108, 111})
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}
