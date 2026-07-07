package ascii

import (
	"fmt"
	"os"
	"strings"
)


func LoadBanner(bannerfile string) (map[rune][]string, error) {
	data, err := os.ReadFile(bannerfile)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	if len(lines) != 856 {
		return nil, fmt.Errorf("bannerFile incomplete")
	}

	banner := map[rune][]string{}

	for i := ' '; i <= '~'; i++ {
		index := int(i-32)
		startpoint := index * 9
		endpoint := 9

		banner[rune(i)] = lines[startpoint+1:startpoint+endpoint]
	}

	return banner, nil
}