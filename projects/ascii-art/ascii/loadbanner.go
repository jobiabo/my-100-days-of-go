package ascii

import (
	"os"
	"strings"
)

// func LoadBanner(filename string) (map[rune][]string, error) {
// 	file, err := os.ReadFile(filename)
// 	if err != nil {
// 		return nil, err
// 	}
// 	//data := strings.ReplaceAll(string(file), "\r\n", "\n")
// 	lines := strings.Split(string(file), "\n")

// 	banner := map[rune][]string{}

// 	if len(lines) != 855 {
// 		return nil, fmt.Errorf("invalid banner file")
// 	}

// 	for i := 32; i <= 126; i++ {
// 		index := rune(i - 32)
// 		startpoint := i*9 + 1
// 		endpoint := startpoint + 9

// 		line := lines[startpoint:endpoint]
// 		banner[index] = line
// 	}
// 	return banner, nil
// }

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
