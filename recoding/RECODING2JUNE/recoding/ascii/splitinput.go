package ascii

import "strings"

func SplitInput(input string) []string{
	return strings.Split(input, "\\n")
}