package main

import (
	"html/template"
	"net/http"

	"github.com/madzumo/redmineapi/internal"
)

// handlers
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "GO")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/nav.tmpl",
		"./ui/html/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	// err = ts.Execute(w, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// w.Write([]byte("Hello from Snippetbox"))
}

func (app *application) sendTicketPost(w http.ResponseWriter, r *http.Request) {
	//this is not needed because we are handling on the servemux (router)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusCreated) //201

	internal.RedmineTicket("sub", "desc")
	w.Write([]byte("Save new snippet..."))
}
