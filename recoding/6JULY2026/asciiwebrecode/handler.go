package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
	Text   string
	Banner string
}

var tmpl = template.Must(template.ParseFiles("index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, nil)
}

func AsciiHandlder(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	if r.URL.Path != "/ascii" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	art, err := Ascii(text, banner)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data.Result = art
	data.Text = text
	data.Banner = banner

	tmpl.Execute(w, data)

}

func AsciiswitchHandlder(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	if r.URL.Path != "/ascii-switch" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")

	art, err := Ascii(text, banner)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data.Result = art
	data.Text = text
	data.Banner = banner

	tmpl.Execute(w, data)

}
