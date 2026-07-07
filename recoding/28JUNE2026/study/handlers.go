package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Input  string
	Banner string
	Result string
}

var templ = template.Must(template.ParseFiles("index1.html"))
var tmpl = template.Must(template.ParseFiles("index2.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}

	templ.Execute(w, data)
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	if r.URL.Path != "/ascii" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}

	if r.Method == http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		banner := r.FormValue("banner")

		result, err := AsciiArt(text, banner)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		data.Result = result
		data.Input = text
		data.Banner = banner
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "ERROR SERVING TEMPLATE", http.StatusInternalServerError)
		fmt.Println("err:= ", err)
	}
}

func AsciiSwitchHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	if r.URL.Path != "/ascii-switch" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

	if r.Method == http.MethodGet {
		text := r.URL.Query().Get("text")
		banner := r.URL.Query().Get("banner")

		result, err := AsciiArt(text, banner)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		log.Println(text)
		log.Println(banner)

		data.Result = result
		data.Input = text
		data.Banner = banner
	}

	tmpl.Execute(w, data)
}
