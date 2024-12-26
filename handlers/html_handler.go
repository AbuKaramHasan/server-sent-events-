package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type to HTML
	w.Header().Set("Content-Type", "text/html")

	// Define the path to your HTML template file
	htmlFilePath := filepath.Join("templates", "index.html")

	// Open the HTML file
	file, err := os.Open(htmlFilePath)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Parse the template
	tmpl, err := template.ParseFiles(htmlFilePath)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	// Prepare data to pass to the template
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Welcome!",
		Message: "Hello, this is a dynamic HTML page rendered with Go.",
	}

	// Render the template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
