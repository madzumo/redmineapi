package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// handlers
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "GO")
	message := r.URL.Query().Get("message")
	dataMSG := struct {
		Message string
	}{
		Message: message,
	}
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

	err = ts.ExecuteTemplate(w, "base", dataMSG)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (app *application) sendTicketPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	t := getTicketMeta()
	app.redmine.RedmineURL = t.redmineURL
	app.redmine.ApiKey = t.apiKey
	app.redmine.Issue.Issue.PriorityID = t.priorityID
	app.redmine.Issue.Issue.ProjectID = t.projectID
	app.redmine.Issue.Issue.Description = r.PostForm.Get("details")
	app.redmine.Issue.Issue.Subject = r.PostForm.Get("subj")
	app.redmine.SendTicket()

	http.Redirect(w, r, "/?message=Ticket submitted successfully!", http.StatusSeeOther)
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

	currentE := getENV()
	err = ts.ExecuteTemplate(w, "base", currentE)
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

	// fmt.Printf("What I got: %s, %s, %s", redmine, pid, apiKey)
	os.Setenv("RED_URL", redmine)
	os.Setenv("RED_PID", pid)
	os.Setenv("RED_APIKEY", apiKey)
	http.Redirect(w, r, "/admin/", http.StatusSeeOther)
}
