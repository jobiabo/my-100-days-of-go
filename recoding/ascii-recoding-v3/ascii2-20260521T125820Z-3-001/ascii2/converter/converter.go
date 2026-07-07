package converter

import (
	"strings"
)

func StringToArt(input string) string {
	if input == "" {
		return ""
	}

	patterns := map[rune][]string{
		'0': {
			" ___ ",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},

		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},

		'2': {
			" ___ ",
			"    |",
			" ___|",
			"|    ",
			"|___ ",
		},

		'3': {
			" ___ ",
			"    |",
			" ___|",
			"    |",
			" ___|",
		},

		'4': {
			"|   |",
			"|   |",
			"|___|",
			"    |",
			"    |",
		},

		'5': {
			" ___ ",
			"|    ",
			"|___ ",
			"    |",
			" ___|",
		},

		'6': {
			" ___ ",
			"|    ",
			"|___ ",
			"|   |",
			"|___|",
		},

		'7': {
			" ___ ",
			"    |",
			"    |",
			"    |",
			"    |",
		},

		'8': {
			" ___ ",
			"|   |",
			"|___|",
			"|   |",
			"|___|",
		},

		'9': {
			" ___ ",
			"|   |",
			"|___|",
			"    |",
			" ___|",
		},
	}

	lines := strings.Split(input, "\n")

	finalOutput := []string{}

	for _, textLine := range lines {

		// create 5 empty rows
		rows := make([]string, 5)

		// loop through each digit
		for _, digit := range textLine {

			// validate digit
			if digit < '0' || digit > '9' {
				return ""
			}

			// append each pattern row
			for row := 0; row < 5; row++ {
				rows[row] += patterns[digit][row]
			}
		}

		// join 5 rows together
		rendered := strings.Join(rows, "\n")

		finalOutput = append(finalOutput, rendered)
	}

	// join all rendered blocks
	return strings.Join(finalOutput, "\n") + "\n"
}
