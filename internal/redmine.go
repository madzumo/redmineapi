package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Issues struct {
	Issue struct {
		ProjectID   string `json:"project_id"`
		Subject     string `json:"subject"`
		Description string `json:"description"`
		PriorityID  string `json:"priority_id"`
	} `json:"issue"`
}

type RedmineTicket struct {
	Issue      Issues
	RedmineURL string
	ApiKey     string
}

func (t *RedmineTicket) SendTicket() {
	issueData, err := json.Marshal(t.Issue)
	if err != nil {
		fmt.Printf("Error marshalling issue: %v\nJSON sent:%s", err, t.Issue)
		return
	}

	req, err := http.NewRequest("POST", t.RedmineURL, bytes.NewBuffer(issueData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Redmine-API-Key", t.ApiKey)
	//create a ticket on behalf of another user
	// req.Header.Set("X-Redmine-Switch-User", "brian") //user_login_or_id

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
