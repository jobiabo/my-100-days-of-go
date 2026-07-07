package converter

import (
	"strings"
)

func StringToArt(input string) string {
	if input == "" {
		return ""
	}
	patterns := map[rune][]string {
		'0': []string{
                " ___ ",
                "|   |",
                "|   |",
                "|   |",
                "|_ _|",
                
	         },

		'1': []string{
                "  |  ",
                "  |  ",
                "  |  ",
                "  |  ",
                "  |  ",
                
	         },

		'2': []string{
                " ___ ",
                "    |",
                " ___|",
                "|    ",
                "| __ ",
                
	         },

		'3': []string{
                " ___ ",
                "    |",
                " ___|",
                "    |",
                " ___|",
                
	         },

		'4': []string{
                "|   |",
                "|   |",
                "|___|",
                "    |",
                "    |",
                
	         },

		'5': []string{
                " ___ ",
                "|    ",
                "|___ ",
                "    |",
                " ___|",
                
	         },

		'6': []string{
                " ___ ",
                "|    ",
                "|___ ",
                "|   |",
                "| __|",
                
	         },

		'7': []string{
                " ___ ",
                "    |",
                "    |",
                "    |",
                "    |",
                
	         },

		'8': []string{
                " ___ ",
                "|   |",
                "|___|",
                "|   |",
                "|___|",
                
	         },

		'9': []string{
                " ___ ",
                "|   |",
                "|___|",
                "    |",
                " ___|",
                
	         },
	}

	splitinput := strings.Split(input, "\n")
	finaloutput := []string{}

	for _, word := range splitinput {
		rows := make([]string, 5)
		for _, char := range word {
			if char < '0' || char > '9' {
				return ""
			}
			for row := 0; row < len(rows); row++ {
				rows[row] += patterns[char][row]
			}
			
		}
		rendered := strings.Join(rows, "\n")
		finaloutput = append(finaloutput, rendered)
	}
	return strings.Join(finaloutput, "\n") + "\n"

}