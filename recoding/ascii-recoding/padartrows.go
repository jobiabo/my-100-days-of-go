package main

import "strings"

func PadArtRows(rows []string, width int) []string {
	result := []string{}

	for row := 0; row < len(rows); row++ {
		padding := width - len(rows[row])

		if padding > 0 {
			result = append(result,
				rows[row]+strings.Repeat(" ", padding))
		} else {
			result = append(result, rows[row])
		}
	}

	return result
}