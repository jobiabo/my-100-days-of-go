package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(IsEven(a))

}

func IsEven(number []int) []int {
	var even []int

	for i := 0; i < len(number); i++ {

		if number[i]%2 == 0 {
			even = append(even, number[i])
		}

	}
	return even

}
