package handlers

import (
	"fmt"
	"mini/channels"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello, World!\n")

	signal := channels.Signal{
		ID:      1,
		Payload: "Hello, from route!",
		Context: r.Context().Done(), // Pass request context for cancellation.
	}

	// Attempt to send the signal to InputChan.
	select {
	case channels.InputChan <- signal:
		fmt.Fprintf(w, "Signal sent to InputChan.\n")
	case <-r.Context().Done():
		fmt.Fprintf(w, "Request has been canceled before sending signal.\n")
		return
	}

	// Wait for a response from OutputChan1 or cancellation.
	select {
	case <-channels.OutputChan1:
		fmt.Fprintf(w, "Signal received from OutputChan1.\n")
	case <-r.Context().Done():
		fmt.Fprintf(w, "Request has been canceled while waiting for OutputChan1.\n")
		return
	}

	// Write a final response.
	w.Write([]byte("Final response written.\n"))

	// If want to write formated use the below:
	// fmt.Fprintf(w, "Signal response based on status.\n")
}
