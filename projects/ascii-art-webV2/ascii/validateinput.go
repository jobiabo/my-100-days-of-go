package ascii

import (
	"fmt"
)

func ValidateInput(input string) (error) {
	for _, char := range input {
		if char < ' ' || char > '~' {
			fmt.Errorf("%v invalid ascii character found", string(char))
		}
	}
	return nil
}