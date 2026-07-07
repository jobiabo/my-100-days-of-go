package main

import (
	"fmt"
	"helper/helper"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println("incomplete commands")
		return
	}

	input := args[1]
	output := args[2]

	data, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	convData := string(data)

	convData = helper.Hex(convData)
	convData = helper.UP(convData)
	convData = helper.Punc(convData)

	err = os.WriteFile(output, []byte(convData), 0644)

	log.Println("file parsed successfully")
}
