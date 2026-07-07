package builder

func ApplyBold(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	result := []string{}

	for _, line := range lines {
		newline := ""
		for _, char := range line {
			newline += string(char) + string(char)
		}
		result = append(result, newline)
	}
	return result
}
