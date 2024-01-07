// File: main.go
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./internal/assets/"))

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/bit/view", bitView)
	mux.HandleFunc("/bit/create", bitCreate)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
