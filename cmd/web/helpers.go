package main

import (
	"fmt"
	"os"
)

type currentENV struct {
	CurrentRedmine string
	CurrentPID     string
	CurrentApiKey  string
}

type ticketMeta struct {
	redmineURL string
	apiKey     string
	projectID  string
	priorityID string
}

func getENV() currentENV {
	s := currentENV{
		CurrentRedmine: os.Getenv("RED_URL"),
		CurrentPID:     os.Getenv("RED_PID"),
		CurrentApiKey:  os.Getenv("RED_APIKEY"),
	}
	return s
}

func getTicketMeta() ticketMeta {
	s := getENV()
	t := ticketMeta{
		redmineURL: fmt.Sprintf("%s/issues.json", s.CurrentRedmine),
		apiKey:     s.CurrentApiKey,
		projectID:  s.CurrentPID,
		priorityID: "",
	}
	return t
}
