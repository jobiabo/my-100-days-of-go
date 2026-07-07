package main

import (
	"log"
	"net/http"
)

func main() {
	// input := "hello World"
	// banner := "banners/standard.txt"

	// result, err := (AsciiArt(input, banner))
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(result)

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/ascii", AsciiHandler)
	mux.HandleFunc("/ascii-switch", AsciiSwitchHandler)

	log.Println("Server running on http://localHost:8080")
	http.ListenAndServe(":8080", mux)
}
