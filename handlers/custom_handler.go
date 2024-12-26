package handlers

import (
	"net/http"
)

func CustomHandler(w http.ResponseWriter, r *http.Request) {
	cw := CustomResponseWriter{}.New(w) // Mimics a static-like constructor.
	cw.Stream("Starting to stream updates...")
}
