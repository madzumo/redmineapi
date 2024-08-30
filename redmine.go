package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	redmineURL = "http://localhost:8080/issues.json"
	apiKey     = "92a96812bb461186a0226b919d4dd11b76871853"
	projectID  = "1"
	priorityID = "2"
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

func main() {
	issue := RedmineIssue{
		Issue: Issue{
			ProjectID:   projectID,
			Subject:     "PC Broken",
			Description: "my computer needs fixing",
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
