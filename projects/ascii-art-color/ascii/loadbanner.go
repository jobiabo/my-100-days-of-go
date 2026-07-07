package ascii

import (
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	cleaned := strings.ReplaceAll(string(data), "\r\n", "\n")

	lines := strings.Split(cleaned, "\n")

	banner := map[rune][]string{}

	for i := 32; i <= 126; i++ {
		index := (i - 32)
		startpoint := index * 9
		endpoint := 9

		banner[rune(i)] = lines[startpoint+1 : startpoint+endpoint]
	}
	return banner, nil
}
