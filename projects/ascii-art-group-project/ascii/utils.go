package ascii

import "strings"

// ValidateArgs checks if the user provided 1 or 2 arguments
func ValidateArgs(arg []string) bool {
	return len(arg) >= 2 && len(arg) <= 3
}

// PrepareInput converts literal "\n" strings into true line breaks
func PrepareInput(text string) string {
	text = strings.ReplaceAll(text, "\\n", "\n")
	text = strings.ReplaceAll(text, "\\t", "    ") // Convert literal tabs to 4 clean spaces
	return text
}

// isOnlyNewlines checks if the text contains nothing but line breaks
// Inside utils.go, update the evaluation check to be completely type-safe:
func IsOnlyNewlines(s string) bool {
	runes := []rune(s)
	if len(runes) == 0 {
		return false
	}
	for _, r := range runes {
		if r != '\n' {
			return false
		}
	}
	return true
}

// SplitLines breaks the input text into individual sentence lines
func SplitLines(text string) []string {
	return strings.Split(text, "\n")
}