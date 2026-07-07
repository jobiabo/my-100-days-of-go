// EXERCISE 5

package main

import (
	"os"
	"fmt"
)

func main() {
	output, err := ReadFile("readfile.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
	
}

func ReadFile(path string) ([]byte, error) {
	bytecontent, err := os.ReadFile(path) 
	if err != nil {
		return nil, err
	}
	return bytecontent, nil
}