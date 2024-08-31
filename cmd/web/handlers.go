package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// handlers
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Pepa Sucia GO")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/nav.tmpl",
		"./ui/html/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	// err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// w.Write([]byte("Hello from Snippetbox"))
}

func snippetVIEW(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// msg := fmt.Sprintf("Display snippet with ID: %d...", id)
	// w.Write([]byte(msg))

	fmt.Fprintf(w, "Displacy snippet with ID: %d...", id)
}

func snippetCREATE(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet......"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	//this is not needed because we are handling on the servemux (router)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusCreated) //201
	w.Write([]byte("Save new snippet..."))
}
