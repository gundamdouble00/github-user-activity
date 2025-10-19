package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Activity struct {
	ID        string      `json:"id"`
	Type      string      `json:"type"`
	Actor     Actor       `json:"actor"`
	Repo      Repo        `json:"repo"`
	Payload   interface{} `json:"payload"`
	Public    bool        `json:"public"`
	CreatedAt time.Time   `json:"created_at"`
}
type Actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}
type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	var userName string
	fmt.Print("github-user-activity: ")
	fmt.Scan(&userName)

	// Create HTTP client
	client := &http.Client{
		Timeout: time.Second * 12,
	}

	URL := "https://api.github.com/users/" + userName + "/events"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-Github-Api-Version", "2022-11-28")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Request fails: %v", res.StatusCode)
	}

	var activities []Activity
	if err = json.NewDecoder(res.Body).Decode(&activities); err != nil {
		log.Fatalf("%v", err)
	}

	for i, activity := range activities {
		fmt.Printf("\n--- Activity %v ---\n", i)
		fmt.Printf("a) Type of activity: %v\n", activity.Type)
		fmt.Printf("b) Actor: %v\n", activity.Actor.Login)
		fmt.Printf("c) URL of github accounnt: %v\n", activity.Actor.URL)
		fmt.Printf("d) Repository: %v\n", activity.Repo.Name)
		fmt.Printf("e) URL of repository: %v\n", activity.Repo.URL)
		fmt.Printf("f) Public: %v\n", activity.Public)
		fmt.Printf("g) Created at: %v\n", activity.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf("h) ID: %v\n", activity.ID)
	}
}
