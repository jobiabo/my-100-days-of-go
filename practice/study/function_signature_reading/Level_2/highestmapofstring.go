package main

import (
	"fmt"
)

func main() {
	mapp := make(map[string]int)
	mapp["one"] = 1
	mapp["two"] = 2
	mapp["three"] = 3
	mapp["four"] = 4
	mapp["five"] = 5

	fmt.Println(HighestMap(mapp))

}

func HighestMap(list map[string]int) string {
	result := ""
	max := make(map[string]int)
	max["zero"] = 0
	for key, v := range list {
		for _, b := range max {
			if v > b {
				max = list
				result = key
			}

		}
	}
	return string(result)
}
