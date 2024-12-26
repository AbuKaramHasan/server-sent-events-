package handlers

import (
	"fmt"
	"net/http"
)

func SimpleHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type to text/event-stream for streaming updates
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Get the flusher to flush data to the client in real-time
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	// Send an initial message to the client
	fmt.Fprintf(w, "data: Starting to stream updates...\n\n")
	flusher.Flush()
}
