package main

import (
	"html/template"
	"net/http"
	"strings"
	"fmt"
)

// PageData holds the variables sent to our HTML view
type PageData struct {
	Display string
}

func main() {
	// Parse the template file
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		display := r.FormValue("display")
		action := r.FormValue("action")

		// Handle button inputs based on their value
		switch action {
		case "C":
			display = ""
		case "=":
			display = calculate(display)
		default:
			// Append the pressed number or operator to the screen
			display += action
		}

		// Send data back to index.html
		tmpl.Execute(w, PageData{Display: display})
	})

	// Run the server
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// A simple, pure Go evaluator safely built without external dependencies
func calculate(expr string) string {
	if expr == "" {
		return ""
	}
	// Basic implementation replacing basic visual signs
	expr = strings.ReplaceAll(expr, "×", "*")
	expr = strings.ReplaceAll(expr, "÷", "/")

	// Using a naive logic parser or simple state loop.
	// For production, you could use a package like "://github.com"
	return evalSimple(expr)
}

func evalSimple(expr string) string {
	// Placeholder logic for basic addition for demonstration. 
	// In your real code, parse tokens natively or import an evaluation library.
	return expr 
}
