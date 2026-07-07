package main

import "fmt"

func main() {
	input := []int{3, 2, 2}
	fmt.Println(findErrorNums(input))
}

func findErrorNums(nums []int) []int {
	count := make(map[int]int)
	missing := 0
	duplicate := 0
	for _, value := range nums {
		count[value]++

		for i := 1; i <= len(nums); i++ {
			if count[i] == 2 {
				duplicate = i
			}
			if count[i] == 0 {
				missing = i
			}
		}
	}
	return []int{duplicate, missing}
}
