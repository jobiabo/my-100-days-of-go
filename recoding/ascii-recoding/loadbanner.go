package main

import (
	"strings"
	"fmt"
	"os"
)

func Loadbanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}	

	Lines := strings.Split(string(data), "\n")
	if len(Lines) != 855 {
		return nil, fmt.Errorf("invalid banner file")
	}
	banner := map[rune][]string{}

	
		for i := 32; i <= 126; i++ {
			index := (i-32)
			startpoint := index * 9
			endpoint := 9
			banner[rune(i)] = Lines[startpoint+1:startpoint+endpoint]
		}
	return banner, nil
}