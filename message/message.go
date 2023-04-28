package message

import "github.com/google/uuid"

type Message struct {
	Content string    `json:"content,omitempty"`
	User    uuid.UUID `json:"user,omitempty"`
}
