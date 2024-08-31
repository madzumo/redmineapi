package main

import (
	"log"
	"net/http"
)

func main() {
	// router - serveMux
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetVIEW)
	mux.HandleFunc("GET /snippet/create", snippetCREATE)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	// web server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
