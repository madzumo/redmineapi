package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Issue struct {
	ProjectID   string `json:"project_id"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
	PriorityID  string `json:"priority_id"`
	// AssignedID  string `json:"assigned_to_id"`
}

type RedmineIssue struct {
	Issue Issue `json:"issue"`
}

func (t *RedmineIssue) SendTicket(redmineURL, apiKey, userID string) {
	// issue := RedmineIssue{
	// 	Issue: Issue{
	// 		ProjectID:   ticket.ProjectID,
	// 		Subject:     ticketSubject,
	// 		Description: ticketDescription,
	// 		PriorityID:  ticket.PriorityID,
	// 		// AssignedID:  assignedID,
	// 	},
	// }

	issueData, err := json.Marshal(t)
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
	req.Header.Set("X-Redmine-Switch-User", userID) //user_login_or_id
	// This allows you to create a ticket on behalf of another user

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
