package main

import (

	"strings"
)

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
	ab.text += text
	return ab
}

func (ab *ArtBuilder) SetStyle(style string) *ArtBuilder {
	validstyles := map[string]bool{
		"normal":  true,
		"bold":    true,
		"italic":  true,
		"outline": true,
	}

	if !validstyles[style] {
		panic("invalid style")
	}
	ab.style = style
	return ab
}

func GenerateArt(c rune) []string {
	result := []string{}

	for i := 0; i < 8; i++ {
		result = append(result, strings.Repeat(string(c), 4))
	}
	return result
}

func (ab *ArtBuilder) Build() string {
	if ab.text == "" {
		return ""
	}

	output := make([]string, 8)

	for _, text := range ab.text {
		charArt := GenerateArt(text)
		for i := 0; i < len(output); i++ {
			output[i] += charArt[i]
		}

	}

	switch ab.style {
	case "italic":
		output = ApplyItalic(output)
	case "bold":
		output = ApplyBold(output)
	case "outline":
		output = ApplyOutline(output)

	}
	return strings.Join(output, "\n")

}

func ApplyItalic(s []string) []string {
	result := []string{}

	for i, row := range s {

		result = append(result, strings.Repeat(" ", i)+row)

	}
	return result
}

func ApplyBold(s []string) []string {
	result := []string{}

	for _, value := range s {
		newLine := ""
		for _, char := range value {
			newLine += string(char) + string(char)
		}
		result = append(result, newLine)
	}
	return result
}

// 

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
