package helper

import (
	"fmt"
	"strings"
)

func UP(input string) string {
	result := []string{}
	spplit := strings.Fields(input)

	for i := 0; i < len(spplit); i++ {
		if spplit[i] == "(up)" {

			spplit[i-1] = strings.ToUpper(spplit[i-1])
			fmt.Println(spplit[i-1])
			result = append(spplit[:i], spplit[i+1:]...)
		}

	}
	return strings.Join(result, " ")
}
