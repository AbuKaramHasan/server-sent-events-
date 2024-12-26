package handlers

import (
	"fmt"
	"net/http"
)

type CustomResponseWriter struct {
	http.ResponseWriter
	flusher http.Flusher
}

// New creates a CustomResponseWriter with defaults.
// This uses a receiver to set up a CustomResponseWriter conveniently.
func (CustomResponseWriter) New(w http.ResponseWriter) CustomResponseWriter {
	cw := CustomResponseWriter{
		ResponseWriter: w,
	}
	if flusher, ok := w.(http.Flusher); ok {
		cw.flusher = flusher
	}
	return cw
}

func (cw *CustomResponseWriter) WriteHeader(statusCode int) {
	cw.ResponseWriter.WriteHeader(statusCode)
}

func (cw *CustomResponseWriter) Write(data []byte) (int, error) {
	cw.flusher.Flush()
	return cw.ResponseWriter.Write(data)

}

func (cw *CustomResponseWriter) Flush() {
	if cw.flusher != nil {
		cw.flusher.Flush()
	}
}

func (cw *CustomResponseWriter) Stream(data string) {
	// Set CORS headers
	cw.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	cw.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	cw.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	cw.ResponseWriter.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	
	// Set the Content-Type to text/event-stream for streaming updates
	cw.ResponseWriter.Header().Set("Content-Type", "text/event-stream")
	cw.ResponseWriter.Header().Set("Cache-Control", "no-cache")
	cw.ResponseWriter.Header().Set("Connection", "keep-alive")

	// Write the data in the SSE (Server-Sent Events) format
	fmt.Fprintf(cw.ResponseWriter, "data: %s\n\n", data)
	if cw.flusher != nil {
		cw.flusher.Flush()
	}
}

