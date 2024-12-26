package handlers

import (
	"net/http"
)

func CustomHandler(w http.ResponseWriter, r *http.Request) {
	cw := NewCustomResponseWrite(w)
	cw.Stream("Starting to stream updates...")
}
