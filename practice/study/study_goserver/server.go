package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("server running at localhost port http://localhost8080")
	http.HandleFunc("/", home)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "This is my first Go server")
	fmt.Fprint(w, "Say hello to my git.....")
}
