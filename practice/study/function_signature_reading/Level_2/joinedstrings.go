package main

import "fmt"

func main() {
	a := "I am Mr Joel"
	b := "and you?"
	fmt.Println(Joinedstrings(a, b))

}

func Joinedstrings(text string, word string) string {
	return text + " " + word
}
