package converter

import "strings"

func StringToArt(input string) string {
	patterns := map[rune][]string {
		'0': {
			" ___ ",
			"|   |",
			"|   |",
			"|   |",
			"| _ |",
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
			"| __|",
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
			"| __|",
		},

		'6': {
			" ___ ",
			"|    ",
			"|___ ",
			"|   |",
			"| __|",
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
			"| __|",
		},

		'9': {
			" ___ ",
			"|   |",
			"|___|",
			"    |",
			"  __|",
		},

	}

	if input == "" {
		return ""
	}

	finaloutput := []string{}

	entry := strings.Split(input, "\n")
	

	for _, text := range entry {
		rows := make([]string, 5)
		for _, char := range text {
			if char < '0' || char > '9' {
				return ""
			}
			for row := 0; row < len(rows); row++ {
				rows[row] += patterns[char][row]
			}
			
		}
		renderred := strings.Join(rows, "\n")
		finaloutput = append(finaloutput, renderred)
	}
	return strings.Join(finaloutput, "\n") + "\n"
}
