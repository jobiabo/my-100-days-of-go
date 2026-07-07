package ascii

import (
	"strings"
)

func MatchString(input string, subinput string) []int {
	position := []int{}

	offset := 0

	for {
		i := strings.Index(input[offset:], subinput)
		if i == -1 {
			break
		}

		start := offset + i
		for j := 0; j < len(subinput); j++ {
			position = append(position, start+j)
			// fmt.Println(start + j)
		}

		offset = start + len(subinput)
	}
	return position
}
