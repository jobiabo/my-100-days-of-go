package ascii 

import (
	"fmt"
)

func ValidateInput(input string) (rune, error) {
	for _, char := range input {
		if char < ' ' || char > '~' {
			return char, fmt.Errorf("charater %s non-ascii character", char)
		}
	}
	return 0, nil
}