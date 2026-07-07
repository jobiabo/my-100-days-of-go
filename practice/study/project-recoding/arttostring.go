package main

import "strings"

func ArtToString(input string) string {
	splitInput := strings.Split(input, "\n")

	for i := 0; i < len(splitInput); i++ {
		if i <= 0 && splitInput[i+1] == " ___ " && splitInput[i+2] == "|   |" && splitInput[i+3] == "|   |" && splitInput[i+4] == "|   |" && splitInput[i+5] == "|___|" {
			return "1"
		}
	}
	return ""
}
