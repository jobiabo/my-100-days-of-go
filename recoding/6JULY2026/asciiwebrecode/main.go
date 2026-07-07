package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/ascii", AsciiHandlder)
	mux.HandleFunc("/ascii-switch", AsciiswitchHandlder)
	log.Println("server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
