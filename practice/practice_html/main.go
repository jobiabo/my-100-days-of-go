package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	// temp := template.Must(template.ParseFiles("index.html"))

	// temp.Execute(w, temp)
	// fmt.Fprint(w, "hello")

	temp, _ := template.ParseFiles("index.html")
	temp.Execute(w, temp)
}

func SecondPage(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("index2.html")
	temp.Execute(w, temp)
}

func ThirdPage(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("index3.html")
	temp.Execute(w, temp)
}

func FouthPage(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("index4.html")
	temp.Execute(w, temp)
}

func main() {
	http.HandleFunc("/home", HomePage)

	http.HandleFunc("/page-two", SecondPage)

	http.HandleFunc("/page-three", ThirdPage)

	http.HandleFunc("/fourth-page", FouthPage)

	fmt.Println("server running on https://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
