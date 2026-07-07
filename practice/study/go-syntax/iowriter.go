package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(count("hello"))
	
}

func count(s strings.Builder) string {
	fmt.Println(s)
    var b strings.Builder
    fmt.Fprint(&s, "Hello")
    return b.String()
}