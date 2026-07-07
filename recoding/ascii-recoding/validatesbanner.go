package main

import (
	"fmt"
)


func ValidateBanner(banner map[rune][]string) error {
	
	if banner == nil {
		return fmt.Errorf("banner is nil")
	}
	if len(banner) != 95 {
		return fmt.Errorf("missing character: ' ' (ASCII 32)")
	}
	for r := 32; r <= 126; r++ {
		
			if len(banner[rune(r)]) != 8 {
				return fmt.Errorf("character %q has %d lines, expected 8", r, len(banner[rune(r)]))
			}
			if len(banner[rune(r)]) == 6 {
				return fmt.Errorf("only six lines here not eight")
			}
			if len(banner[rune(r)]) == 0 {
				fmt.Errorf("error incomplete character")
			}
		
	}
	return nil
	
}
