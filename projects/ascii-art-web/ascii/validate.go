package ascii

import "fmt"

func ValidateInput(input string) (rune, error) {
	for _, char := range input {
		if char < 32 || char > 126 {
			return char, fmt.Errorf("invalid character")
		}
	}
	return 0, nil
}
