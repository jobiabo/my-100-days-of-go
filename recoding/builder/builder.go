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

func (ab *ArtBuilder) Build() string {

	if ab.text == "" {
		return ""
	}

	result := make([]string, 8)

	for _, value := range ab.text {
		charArt := GenerateArt(value)
		for row := 0; row < 8; row++ {
			result[row] = charArt[row]
		}

	}

	switch ab.style {
	case "normal":

	case "bold":
		result = ApplyBold(result)

	case "italic":
		result = ApplyItalic(result)
	case "outline":
		result = ApplyOutline(result)
	}

	return strings.Join(result, "\n")

}
