package ascii

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(bannername string) (map[rune][]string, error) {
	// file := "banner/" + bannername

	data, err := os.ReadFile(bannername)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	if len(lines) != 856 {
		fmt.Errorf("invalid banner File read")
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
