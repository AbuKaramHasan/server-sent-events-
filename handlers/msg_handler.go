package handlers

import (
	"html/template"
	"mini/channels"
	"net/http"
	"time"
)

func MsgHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type to HTML
	w.Header().Set("Content-Type", "text/html")

	// Wait for data from OutputChan1
	var signal channels.Signal
	select {
	case signal = <-channels.OutputChan1:
		// Successfully received the signal
	case <-time.After(5 * time.Second): // Timeout after 5 seconds
		http.Error(w, "Timeout waiting for signal", http.StatusGatewayTimeout)
		return
	case <-r.Context().Done(): // Handle request cancellation
		http.Error(w, "Request canceled by client", http.StatusRequestTimeout)
		return
	}

	// Define the path to your HTML template file
	htmlFilePath := "templates/index.html"

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
		Title:   "Signal Received!",
		Message: signal.Payload,
	}

	// Render the template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}
