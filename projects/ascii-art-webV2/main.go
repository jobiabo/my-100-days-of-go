package main

import (
	"ascii/handler"
	"fmt"
	"net/http"
)

// type PageData struct {
// 	Output string
// }

// var templ = template.Must(template.ParseFiles("template/index.html"))

func main() {
	http.HandleFunc("/ascii", handler.HandlerAscii)
	fmt.Println("server running at port http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}

// func HandlerAscii(w http.ResponseWriter, r *http.Request) {
// 	data := PageData{}
// 	log.Println(r.Method)

// 	if r.Method == http.MethodPost {

// 		text := r.FormValue("text")
// 		fonts := r.FormValue("banner")

// 		err := ascii.ValidateInput(text)
// 		if err != nil {
// 			http.Error(w, "", http.StatusNotFound)
// 		}

// 		fmt.Println(text)
// 		fmt.Println(fonts)

// 		font, err := ascii.LoadBanner(fonts)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		asciiart := ascii.Generate(text, font)
// 		data.Output = asciiart

// 		fmt.Println(asciiart)

// 	}

// 	if r.Method == http.MethodGet {

// 		text := r.URL.Query().Get("text")
// 		fonts := r.URL.Query().Get("banner")

// 		err := ascii.ValidateInput(text)
// 		if err != nil {
// 			http.Error(w, "", http.StatusNotFound)
// 		}

// 		log.Println(text)
// 		log.Println(fonts)

// 		font, err := ascii.LoadBanner(fonts)
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		asciiart := ascii.Generate(text, font)
// 		data.Output = asciiart

// 		fmt.Println(asciiart)

// 	}
// 	templ.Execute(w, data)
// }
