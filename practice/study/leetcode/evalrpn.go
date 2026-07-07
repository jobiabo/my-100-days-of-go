package main

import (
	"fmt"
	"strconv"
)

func main() {
	Input := []string{"7", "2", "1", "+", "3", "*"}
	fmt.Println(evallRPN(Input))
}

func evallRPN(tokens []string) int {
	stack := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			// fmt.Println("1st value of stack", len(stack))
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			fmt.Println(b)
			fmt.Println(a)

			stack = stack[:len(stack)-2]
			fmt.Println(stack)
			stack = append(stack, a+b)
			fmt.Println(stack)

		case "-":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)

		case "*":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)

		case "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)

		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)

		}
	}
	return stack[0]

}
