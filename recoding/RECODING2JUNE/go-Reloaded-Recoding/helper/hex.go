package helper

import (
	"fmt"
	"strconv"
	"strings"
)

func Hex(input string) string {
	result := []string{}
	spplit := strings.Fields(input)

	for i := 0; i < len(spplit); i++ {
		if spplit[i] == "(hex)" {
			num, err := strconv.ParseInt(spplit[i-1], 16, 64)
			if err != nil {
				fmt.Println(err)
			}
			// numalp := strconv.FormatInt(num, 10)
			spplit[i-1] = fmt.Sprintf("%d", num)
			fmt.Println(spplit[i-1])
			result = append(spplit[:i], spplit[i+1:]...)
		}

	}
	return strings.Join(result, " ")
}
