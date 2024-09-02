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
	// issue2 := RedmineIssue{
	// 	Issue: Issue{
	// 		ProjectID:   "22",
	// 		Subject:     "33",
	// 		Description: "44",
	// 		PriorityID:  "55",
	// 		// AssignedID:  assignedID,
	// 	},
	// }

	// fmt.Printf("t: %s\n", t)
	// fmt.Printf("issue2: %s\n", issue2)
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
	// req.Header.Set("X-Redmine-Switch-User", "user_login_or_id") //create a ticket on behalf of another user

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
