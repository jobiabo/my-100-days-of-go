package main

import "fmt"

func main() {
	a := []int{1, 3}
	fmt.Println(buildArray(a, 3))

}

func buildArray(target []int, n int) []string {

	stack := []string{}

	j := 0
	for i := 1; i <= n && j < len(target); i++ {
		stack = append(stack, "Push")
		

		if i == target[j] {
			j++
		} else {
			stack = append(stack, "Pop")
		}
		
	}
	return stack

}
