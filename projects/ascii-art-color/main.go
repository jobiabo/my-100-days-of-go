package main

import (
	"ascii-art-color/ascii"
	"fmt"
)

var cmd = ascii.ParseArguments()

func main() {

	_, err := ascii.ValidateInput(cmd.Text)
	if err != nil {
		fmt.Println(err)
		return
	}
	fonts, err := ascii.LoadBanner("banners/standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	Art := ascii.GenerateArt(cmd.Text, fonts)
	fmt.Println(Art)

}
