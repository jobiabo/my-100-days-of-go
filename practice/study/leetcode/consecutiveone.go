package main

import "fmt"

func main() {
	Input := []int{1,1,0,1,1,1}
	fmt.Println(findMaxConsecutiveOnes(Input))
}

func findMaxConsecutiveOnes(nums []int) int {
	result := 0
	max := 0
	for i := 0; i < len(nums); i++ {

		if nums[i] == 1 {
			max++
			if max > result {
				result = max
			}
		} else {
			max = 0
		}

	}

	
	return result

}
