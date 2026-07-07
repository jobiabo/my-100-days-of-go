package main

import "fmt"

func main() {
	input := []int{2, 5, 1, 3, 4, 7}
	output := []int{}
	x := []int{}
	y := []int{}

	for j := 0; j < len(input)/2; j++ {
		x = append(x, input[j])
	}
	for k := len(input) / 2; k < len(input); k++ {
		y = append(y, input[k])
	}

	for i := 0; i < len(x); i++ {
		output = append(output, x[i])
	
		output = append(output, y[i])

	}

	// fmt.Println(x)
	// fmt.Println(y)
	fmt.Println(output)

}
