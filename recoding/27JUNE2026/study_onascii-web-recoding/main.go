package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandlers)
	mux.HandleFunc("/ascii", asciiHandler)
	mux.HandleFunc("/ascii-switch", asciiHandlerSwitch)
	log.Println("server runing at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
