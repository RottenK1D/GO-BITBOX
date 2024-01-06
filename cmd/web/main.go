// File: main.go
package main

import (
	"log"
	"net/http"

	"github.com/RottenK1D/GO-BITBOX/ui/html/pages"
	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/bit/view", bitView)
	mux.HandleFunc("/bit/create", bitCreate)
	mux.Handle("/hello", templ.Handler(pages.Hello()))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
