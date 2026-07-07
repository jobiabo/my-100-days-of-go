package handler

import (
	"log"
	"net/http"
	"html/template"
	"fmt"
	// "ascii/ascii"
	// "strings"
)




var templ = template.Must(template.ParseFiles("templates/index.html"))
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	if r.URL.Path != "/" {
		fmt.Fprint(w, "route not accepted")
		log.Println(r.URL.Path)
		http.Error(w, "Page Not Found", http.StatusNotFound)
	 

	}

	templ.Execute(w, data)

}