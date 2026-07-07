package main

import (
	"fmt"
	"os"
	"strings"
)

func Ascii(input string, banner string) (string, error) {
	data, err := os.ReadFile(banner + ".txt")
	if err != nil {
		return "error from load", err
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	if len(lines) != 856 {
		return "incomplete banner", fmt.Errorf("banner error")
	}

	bannerMap := map[rune][]string{}

	for i := 32; i <= 126; i++ {
		index := i - 32
		startpoint := index * 9
		endpoint := 9
		bannerMap[rune(i)] = lines[startpoint+1 : startpoint+endpoint]
	}

	return strings.Join(bannerMap[rune(65)], "\n"), nil
}
