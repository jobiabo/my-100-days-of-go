package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	intersectedValues := []int{}
	charMap := map[int]int{}
	for _, char := range nums1 {
		charMap[char]++
	}
	fmt.Println(charMap)

	for _, ch := range nums2 {
		_, exist := charMap[ch]
		if exist {
			intersectedValues = append(intersectedValues, ch)
		}

	}
	return intersectedValues
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersection(nums1, nums2))
}
