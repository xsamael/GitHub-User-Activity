package model

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func ReqGithub(name string) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", name)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var event []Event

	err = json.Unmarshal(content, &event)
	if err != nil {
		fmt.Println("Error al decodificar JSON:", err)
		return
	}

	for _, event := range event {
		switch event.Type {
		case "PushEvent":
			var payload PushPayload
			if err := json.Unmarshal(event.Payload, &payload); err != nil {
				fmt.Printf("Pushed %d commits to %s\n", len(payload.Commits), event.Repo.Name)
			}
		case "IssuesEvent":
			fmt.Printf("Opened a new issue in %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("Starred %s\n", event.Repo.Name)
		case "ForkEvent":
			fmt.Printf("Forked %s\n", event.Repo.Name)
		default:
			fmt.Printf("- %s in %s\n", strings.TrimSuffix(event.Type, "Event"), event.Repo.Name)
		}
	}
}
