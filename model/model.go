package model

import "encoding/json"

type Event struct {
	Type      string          `json:"type"`
	Repo      Repo            `json:"repo"`
	Payload   json.RawMessage `json:"payload"`
	CreatedAt string          `json:"created_at"`
}

type Repo struct {
	Name string `json:"name"`
}

type PushPayload struct {
	Commits []Comit `json:"commits"`
}

type Comit struct {
	Message string `json:"message"`
}
