package main
import (
	"strings"
	"fmt"
)

func main() {
	// a := []string{"HELLO", "HELLO","HELLO","HELLO","HELLO","HELLO","HELLO",}
	// b := ApplyOutline(a)
	// b := ApplyItalic(a)
	// b := ApplyBold(a)
	b := GenerateArt('a')


	fmt.Println(strings.Join(b, "\n"))
}