package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"study/ascii"
)

func main() {
	if !ascii.ValidateArgs(os.Args) {
		fmt.Fprintf(os.Stderr, "Usage: go run . [STRING][BANNER]\n")
		fmt.Fprintf(os.Stderr, "Example: go run . \"hello\"standard\n")
		os.Exit(1)
	}

	inputText := os.Args[1]
	bannerName := "standard"
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}

	bannerName = strings.TrimSuffix(bannerName, ".txt")

	bannerFile := filepath.Clean(bannerName + (".txt"))

	ProcessedText := ascii.PrepareInput(inputText)
	if ascii.IsOnlyNewlines(ProcessedText) {
		fmt.Print(strings.Repeat("\n", len(ProcessedText)))
		return
	}

	banner, err := ascii.LoadBanner("banners/" + bannerFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: could not load bannerfile '%s': %v\n", bannerFile, err)
		os.Exit(2)
	}

	splitInput := ascii.SplitLines(ascii.PrepareInput(inputText))

	result := ascii.RenderText(splitInput, banner)
	fmt.Print(result)

}
