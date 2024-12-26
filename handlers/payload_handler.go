package handlers

import (
	"net/http"
)

func PayloadHandler(w http.ResponseWriter, r *http.Request) {
	cw := CustomResponseWriter{}.New(w) // Create a new custom writer.

	// Create a payload and stream it
	payload := Payload{
		Event:   "start",
		Message: "Starting to stream updates...",
	}
	cw.StreamPayload(payload)
}
