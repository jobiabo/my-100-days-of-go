package main

import (
	"fmt"
)
func main() {
	a := []int{7,7,7,7}
	fmt.Println(smallerNumbersThanCurrent(a))
	
}

func smallerNumbersThanCurrent(nums []int) []int {
	result := []int{}
	



	for i:= 0; i < len(nums); i++ {
		others := []int{}
		for j := 0; j < len(nums); j++ {
			
			if nums[j] < nums[i] {
				others = append(others, nums[j]) 
			}
		}

		result = append(result, len(others))


	}
	return result
}