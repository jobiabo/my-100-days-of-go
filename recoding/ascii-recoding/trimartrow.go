package main

import "strings"

func TrimArtRows(rows []string) []string {
	result := make([]string, len(rows))

	for index, value := range rows {
		result[index] = strings.TrimRight(value, " ")
	}
	return result
}
