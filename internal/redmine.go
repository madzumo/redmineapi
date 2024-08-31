package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var (
	redmineURL = "http://localhost:8080/issues.json"
	apiKey     = "1"
	projectID  = "1"
	priorityID = "1"
)

type Issue struct {
	ProjectID   string `json:"project_id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
	PriorityID  string `json:"priority_id"`
}

type RedmineIssue struct {
	Issue Issue `json:"issue"`
}

func RedmineTicket(ticketSubject, ticketDescription string) {
	getENV()
	issue := RedmineIssue{
		Issue: Issue{
			ProjectID:   projectID,
			Subject:     ticketSubject,
			Description: ticketDescription,
			PriorityID:  priorityID,
		},
	}

	issueData, err := json.Marshal(issue)
	if err != nil {
		fmt.Printf("Error marshalling issue: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", redmineURL, bytes.NewBuffer(issueData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Redmine-API-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Issue successfully created!")
	} else {
		fmt.Printf("Failed to create issue. Status: %s\n", resp.Status)
	}
}

func getENV() {
	redmineURL = fmt.Sprintf("%s/issues.json", os.Getenv("RED_URL"))
	apiKey = os.Getenv("RED_APIKEY")
	projectID = os.Getenv("RED_PROJECTID")
}
