package handler

import (
	"net/http"
	// "html/template"
	"log"
	"fmt"
	"ascii/ascii"
)


type PageData struct {
	Result string
	Text string
	Banner string
}
// var tmpl = template.Must(template.ParseFiles("templates/switcher.html"))
func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	if r.URL.Path != "/ascii" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "not alowed", http.StatusMethodNotAllowed)
		return
	}


	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		_, err := ascii.ValidateInput(text) 
		if err != nil {
			http.Error(w, "character should contain only a-z/A-Z/0-9 & punctuations", http.StatusBadRequest)
			return
		}

		fonts, err := ascii.LoadBanner(banner)
		if err != nil {
			log.Println(err)
			http.Error(w, "error from the server",http.StatusInternalServerError)
			return
		}
		data.Result = ascii.GenerateArt(text, fonts)
		data.Text = text
		data.Banner = banner
		fmt.Println(data.Result)
	}

	tmpl.Execute(w, data)
	
}