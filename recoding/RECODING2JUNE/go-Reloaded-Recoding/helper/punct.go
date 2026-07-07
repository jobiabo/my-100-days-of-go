package helper

import "regexp"

func Punc(input string) string {
	re1 := regexp.MustCompile(`\s([!])`)
	input = re1.ReplaceAllString(input, "$1")

	re2 := regexp.MustCompile(`'\s(.*?)\s'`)
	input = re2.ReplaceAllString(input, "'$1'")
	return input
}
