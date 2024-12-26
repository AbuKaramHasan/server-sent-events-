package handlers

import (
	"fmt"
	"html/template"
	"mini/channels"
	"net/http"
	"time"
)

// HTMLHandler serves the initial HTML page and streams updates to the client.
func SSEHandler(w http.ResponseWriter, r *http.Request) {
	// Check if this is a request for the EventStream
	if r.URL.Query().Get("stream") == "true" {
		handleStream(w, r)
		return
	}

	// Serve the initial HTML template
	w.Header().Set("Content-Type", "text/html")

	// Load the HTML template (event.html)
	tmpl, err := template.ParseFiles("templates/event.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// Render the template with initial data
	data := struct {
		Title string
	}{
		Title: "Real-Time Updates",
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

// handleStream handles real-time updates via Server-Sent Events (SSE).
func handleStream(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Get the flusher for streaming updates
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Stream updates from the InputChan
	for {
		select {
		case signal := <-channels.OutputChan1:
			// Stream the new signal payload to the client
			fmt.Fprintf(w, "data: %s\n\n", signal.Payload)
			flusher.Flush()
		case <-time.After(30 * time.Second):
			// Send a heartbeat to keep the connection alive
			fmt.Fprintf(w, "data: heartbeat\n\n")
			flusher.Flush()
		case <-r.Context().Done():
			// Client disconnected or request was canceled
			fmt.Println("Client disconnected from streaming")
			return
		}
	}
}
