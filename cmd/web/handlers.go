package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Home is a simple HTTP handler function which writes a byte slice
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from BITBOX"))
}

// BitView displays a specific bit based on ID provided in the URL
func bitView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific bit with ID %d...", id)
}

// BitCreate creates a new bit
func bitCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method not allowed", 405)
		return
	}

	w.Write([]byte("Create a new bit..."))
}
