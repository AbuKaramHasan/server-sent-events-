package handlers

import (
	"fmt"
	"net/http"
)

type CustomResponseWritter struct {
	http.ResponseWriter
	flusher http.Flusher
}

func NewCustomResponseWrite(w http.ResponseWriter) *CustomResponseWritter {
	if flusher, ok := w.(http.Flusher); !ok {
		fmt.Println("ResponseWriter does not support Flusher interface")
		return &CustomResponseWritter{
			ResponseWriter: w}
	} else {
		return &CustomResponseWritter{
			ResponseWriter: w,
			flusher:        flusher,
		}
	}
}

func (w *CustomResponseWritter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *CustomResponseWritter) Write(data []byte) (int, error) {
	w.flusher.Flush()
	return w.ResponseWriter.Write(data)

}

func (w *CustomResponseWritter) Flush() {
	if w.flusher != nil {
		w.flusher.Flush()
	}
}

func (w *CustomResponseWritter) Stream(data string) {
	// Set the Content-Type to text/event-stream for streaming updates
	w.ResponseWriter.Header().Set("Content-Type", "text/event-stream")
	w.ResponseWriter.Header().Set("Cache-Control", "no-cache")
	w.ResponseWriter.Header().Set("Connection", "keep-alive")
	fmt.Fprintf(w.ResponseWriter, "data: %s\n\n", data)
	if w.flusher != nil {
		w.flusher.Flush()
	}
}
