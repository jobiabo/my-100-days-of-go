// EXERCISE 16
package main

import (
	"fmt"
	"strings"
)

func main() {
	textinput := "hello world"
	fmt.Println(SplitTrans(textinput))

}

func SplitTrans(data string) string {
	splitdata := strings.Fields(data)
	for i := 0; i < len(splitdata); i++ {
		splitdata[i] = strings.ToUpper(splitdata[i])
	}
	return strings.Join(splitdata, "-")
	// all steps were done and manipulated on the default datatype 'string'
}
