package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Result  string
	Input   string
	Banners string
}

var templ = template.Must(template.ParseFiles("index.html"))
var tmpl = template.Must(template.ParseFiles("index2.html"))

func HomeHandlers(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "not found", http.StatusNotFound)
	}

	templ.Execute(w, nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}

	if r.URL.Path != "/ascii" {
		http.Error(w, "page not found", http.StatusNotFound)
	}

	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		banner := r.FormValue("banner")

		log.Println(text)
		log.Println(banner)

		result, err := AsciiArt(text, banner)
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		data.Result = result
		data.Input = text
		data.Banners = banner

	}
	tmpl.Execute(w, data)

}

func asciiHandlerSwitch(w http.ResponseWriter, r *http.Request) {
	data := PageData{}

	if r.URL.Path != "/ascii-switch" {
		http.Error(w, "page not found", http.StatusNotFound)
	}

	if r.Method == http.MethodGet {
		text := r.URL.Query().Get("text")
		banner := r.URL.Query().Get("banner")

		result, err := AsciiArt(text, banner)
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		data.Result = result
		data.Input = text
		data.Banners = banner

		log.Println("pored:", data.Input, "and goten", text)
		log.Println("pored:", data.Banners, "and goten", banner)

	}
	tmpl.Execute(w, data)

}
