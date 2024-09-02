package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux() // router - serveMux

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("POST /", app.sendTicketPost)
	mux.HandleFunc("GET /admin/", app.adminArea)
	mux.HandleFunc("POST /admin/", app.adminAreaPost)
	return mux
}
