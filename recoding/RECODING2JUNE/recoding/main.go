package main

import (
	"net/http"
	"log"
	"ascii/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/ascii", handler.AsciiHandler)
	mux.HandleFunc("/ascii-switch", handler.AsciiSwitcherHandler)
	log.Println("server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}