package main

import "fmt"

func main() {
	a := []int{1, 1}
	fmt.Println(findDisappearedNumbers(a))

}

func findDisappearedNumbers(nums []int) []int {
	finder := make(map[int]int)
	check := make(map[int]int)
	missingno := []int{}
	for key, value := range nums {
		finder[value] = nums[key]
	}
	for i := 1; i <= len(nums); i++ {
		_, ok := finder[i]
		if !ok {
			check[i] = i
			missingno = append(missingno, i)
		}
	}
	return missingno
}
