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
		// Context: r.Context().Done(), // Pass request context for cancellation.
	}

	// Attempt to send the signal to InputChan.
	select {
	case channels.InputChan <- signal:
		fmt.Fprintf(w, "Signal sent to InputChan.\n")
	case <-r.Context().Done():
		fmt.Fprintf(w, "Request has been canceled before sending signal.\n")
		return
	}

	// Process signals from OutputChan1 or cancellation.
	var result []string
	select {
	case signal := <-channels.OutputChan1:
		result = append(result, fmt.Sprintf("Signal received from OutputChan1: %v", signal.Payload))
	case <-r.Context().Done():
		fmt.Fprintf(w, "Request has been canceled while waiting for OutputChan1.\n")
		return
	}

	// Aggregate results and send them all at once.
	fmt.Fprintf(w, "Final response written.\n")
	fmt.Fprintf(w, "Signals processed: \n")
	for _, line := range result {
		fmt.Fprintf(w, "%s\n", line)
	}
}
