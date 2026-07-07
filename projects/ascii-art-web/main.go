package main

import (
	"ascii/ascii"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type PageData struct {
	AsciiArt string
}

var templ = template.Must(template.ParseFiles("templates/index.html"))

func main() {

	fmt.Println("Server running on http://localhost:8081")

	http.HandleFunc("/ascii", HomeAsciiHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func HomeAsciiHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.Method)

	data := PageData{}

	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		font := r.FormValue("font")

		draft, err := ascii.LoadBanner("banners/" + font)
		if err != nil {
			fmt.Println("LoadBanner error:", err)
		}

		asciiArt := ascii.GenerateArt(text, draft)

		f, err := os.OpenFile("activityLog.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644)
		if err != nil {
			return
		}
		defer f.Close()

		fmt.Fprintf(f, "Method: %s\n", r.Method)
		fmt.Fprintf(f, "Text: %s\n", text)
		fmt.Fprintf(f, "Font: %s\n", font)
		fmt.Fprintf(f, "Generated:\n%s\n", asciiArt)

		data.AsciiArt = asciiArt
	}

	err := templ.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
