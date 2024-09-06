package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Fname string `json:"firstname"`
	Lname string `json:"lastname"`
	Email string `json:"mail"`
}

type UserResponse struct {
	User   []User `json:"users"`
	Tcount string `json:"total_count"`
	Offset string `json:"offset"`
	Limit  string `json:"limit"`
}

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
	UserID     string
	UserEmail  string
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
	// req.Header.Set("X-Redmine-Switch-User", "watson") //user_login_or_id

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

func (t *RedmineTicket) GetUserID() {
	url := fmt.Sprintf("%s/users.json?key=%s&mail=%s", t.RedmineURL, t.ApiKey, t.UserEmail)

	resp, err := http.Get(url)
	if err != nil {
		t.UserID = ""
		fmt.Println("Error reading response in UserID lookup", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body in UserID Lookup", err)
		return
	}

	var uR UserResponse
	err = json.Unmarshal(body, &uR)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if len(uR.User) > 0 {
		t.UserID = uR.User[0].ID
		fmt.Printf("User ID: %s\n", t.UserID)
	} else {
		fmt.Println("No user found with that e-mail")
	}
}
