package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/madzumo/redmineapi/internal"
)

// handlers
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "GO")

	files := []string{
		"./ui/html/base.go.html",
		"./ui/html/nav.go.html",
		"./ui/html/home.go.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) sendTicketPost(w http.ResponseWriter, r *http.Request) {
	//this is not needed because we are handling on the servemux (router)
	// if r.Method != http.MethodPost {
	// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	// 	return
	// }
	w.WriteHeader(http.StatusCreated) //201

	internal.RedmineTicket("sub", "desc")
	w.Write([]byte("Save new snippet..."))
}

func (app *application) adminArea(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.go.html",
		"./ui/html/nav.go.html",
		"./ui/html/admin.go.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) adminAreaPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		return
	}
	redmine := r.PostForm.Get("redmine")
	pid := r.PostForm.Get("pid")
	apiKey := r.PostForm.Get("apikey")

	os.Setenv("RED_URL", redmine)
	os.Setenv("RED_PID", pid)
	os.Setenv("RED_APIKEY", apiKey)
	// fmt.Fprintf(w, "%s", redmine)
}
