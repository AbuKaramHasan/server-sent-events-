package handlers

import (
	"fmt"
	"mini/channels"
	"net/http"
	"time"
)

func EventHandler(w http.ResponseWriter, r *http.Request) {
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

	// Listen for updates on the OutputChan1 in a loop
	for {
		select {
		case signal := <-channels.OutputChan1:
			// Stream the received signal payload to the client
			fmt.Fprintf(w, "data: %s\n\n", signal.Payload)
			flusher.Flush()
		case <-time.After(30 * time.Second):
			// Send a heartbeat every 30 seconds to keep the connection alive
			fmt.Fprintf(w, "data: heartbeat\n\n")
			flusher.Flush()
		case <-r.Context().Done():
			// Handle client disconnection or request cancellation
			fmt.Println("Client disconnected")
			return
		}
	}
}
