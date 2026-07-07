package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println(ListLines("evennumbers.go"))
	g := (ListLines("evennumbers.go"))
	k := strings.Join(g, "\n")
	fmt.Println(k)

}

func ListLines(file string) []string {
	filecontent, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(filecontent), "\n")
	return lines

}
