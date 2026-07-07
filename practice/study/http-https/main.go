package main

import (
	"fmt"
	"log"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A secure Web Server")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", TestHandler)

	log.Println("server running at http://localost:443")
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", mux)
	if err != nil {
		log.Println(err)
	}
}
