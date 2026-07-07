package ascii

import (
	"fmt"
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
	if len(lines) == 0 {
		return nil, fmt.Errorf("banner is empty")
	}
	if len(lines) != 856 {
		return nil, fmt.Errorf("Invalid banner: expected 856 lines, but got %d lines", len(lines))
	}

    banner := map[rune][]string{}

    for i := 32; i <= 126; i++ {
        index := (i - 32)
        startpoint := index * 9
        endpoint := 9

        banner[rune(i)] = lines[startpoint+1 : startpoint+endpoint]
    }
    return banner, nil
}