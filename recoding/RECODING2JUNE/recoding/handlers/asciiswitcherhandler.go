package handler

import (
	"net/http"
	"html/template"
	"log"
	"ascii/ascii"
	"fmt"
	"strings"
)

var tmpl = template.Must(template.ParseFiles("templates/switcher.html"))
func AsciiSwitcherHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	if r.URL.Path != "/ascii-switch" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodPost {
		http.Error(w, "Method NNNNNNNOt allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodGet {

		text := r.URL.Query().Get("text")
		banner := r.URL.Query().Get("banner")
		_, err := ascii.ValidateInput(text) 
		if err != nil {
			http.Error(w, "character should contain only a-z/A-Z/0-9 & punctuations", http.StatusBadRequest)
			return
		}

		if !(strings.HasSuffix(banner, ".txt")) {
			banner = banner+".txt"
			fmt.Println("refined banner:", banner)

		}

		fmt.Println("refined banner:", banner)

		fonts, err := ascii.LoadBanner("banners/"+banner)
		if err != nil {
			fmt.Println("here is banner from get:", banner)
			log.Println(err)
			http.Error(w, "error from the server",http.StatusInternalServerError)
			return
		}
		data.Result = ascii.GenerateArt(text, fonts)
		data.Text = text
		fmt.Println(data.Result)
	}
	tmpl.Execute(w, data)
}