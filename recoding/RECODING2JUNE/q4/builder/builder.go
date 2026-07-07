package builder

import "strings"

type ArtBuilder struct {
	text  string
	style string
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{
		text:  "",
		style: "normal",
	}
}

func (ab *ArtBuilder) AddText(text string) *ArtBuilder {
	ab.text = text
	return ab
}

func (ab *ArtBuilder) SetStyle(style string) *ArtBuilder {
	validstyles := map[string]bool{
		"bold":    true,
		"italic":  true,
		"outline": true,
		"normal":  true,
	}

	if !validstyles[style] {
		panic("invalid style")
	}
	ab.style = style
	return ab
}

func (ab *ArtBuilder) Build() string {
	if ab.text == "" {
		return ""
	}

	result := make([]string, 8)

	for _, line := range ab.text {
		charArt := GenerateArt(line)
		for row := 0; row < 8; row++ {
			result[row] = charArt[row]
		}
	}

	switch ab.style {
	case "italic":
		result = ApplyItalic(result)
		case "bold":
		result = ApplyBold(result)

		case "outline":
		result = ApplyOutline(result)

		case "normal":

	}
	return strings.Join(result, "\n")
}

func ApplyOutline(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	width := len(lines[0])

	boarder := "+" + strings.Repeat("-", width) + "+"

	result := []string{}

	result = append(result, boarder)

	for _, line := range lines {
		result = append(result, "|"+line+"|")
	}
	result = append(result, boarder)

	if len(result) > 8 {
		return result[:8]
	}

	if len(result) < 8 {
		result = append(result, boarder)
	}
	return result
}

func ApplyItalic(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}
	result := []string{}

	for i, line := range lines {
		result = append(result, strings.Repeat(" ", i)+line)
	}

	return result
}

func ApplyBold(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}
	result := []string{}

	for _, line := range lines {
		newline := ""
		for _, char := range line {
			newline += strings.Repeat(string(char), 4)
		}
		result = append(result, newline)
	}
	return result
}

func GenerateArt(c rune) []string {
	result := []string{}

	for row := 0; row < 8; row++ {
		result = append(result, strings.Repeat(string(c), 4))
	}
	return result
}
