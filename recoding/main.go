package main

import (
	"strings"
	"fmt"
)

func main() {
	h := []string{"HELLO","HELLO","HELLO","HELLO","HELLO","HELLO","HELLO",}

	g := ApplyBold(h)

	// art := NewArtBuilder().
	// 	AddText("hello").
	// 	SetStyle("outline").
	// 	Build()

	fmt.Println(strings.Join(g, "\n"))
}
