package main

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(input string, banner string) (string, error) {
	newbanner := "banners/" + banner + ".txt"
	data, err := os.ReadFile(newbanner)
	if err != nil {
		return "banner not read", err
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	if len(lines) != 856 {
		return "incomplete banner", fmt.Errorf("incomplete banners %d lines found", len(lines))
	}

	bannerMap := map[rune][]string{}

	for i := 32; i < 126; i++ {
		index := i - 32
		startpoint := index * 9
		endpoint := 9

		bannerMap[rune(i)] = lines[startpoint+1 : startpoint+endpoint]
	}

	result := []string{}

	for row := 0; row < 8; row++ {
		var draw strings.Builder
		for _, char := range input {
			draw.WriteString(bannerMap[char][row])
		}
		result = append(result, draw.String())
	}

	return strings.Join(result, "\n"), nil
}
