package ascii

func RenderText(lines []string, bannerMap map[rune][]string) string {
	result := ""
	for _, line := range lines {
		if line == "" {
			result += "\n"
			continue
		}
		result += renderLine(line, bannerMap)
	}

	return result
}

func renderLine(line string, bannerMap map[rune][]string) string {
	result := ""
	for row := 0; row < 8; row++ {
		
		for _, char := range line {
			result += bannerMap[char][row]
		}
		result += "\n"
	}
	return result
}
