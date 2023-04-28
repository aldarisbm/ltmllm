package shared

import (
	"github.com/google/uuid"
	"time"
)

type Conversation struct {
	Messages   []*Message `json:"messages,omitempty"`
	CreateTime time.Time  `json:"createTime,omitempty"`
	UpdateTime time.Time  `json:"updateTime,omitempty"`
	Id         uuid.UUID  `json:"id,omitempty"`
}

type Message struct {
	Content    string    `json:"content,omitempty"`
	Id         uuid.UUID `json:"id,omitempty"`
	User       string    `json:"user,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
}

func NewConversation() Conversation {
	return Conversation{
		Messages:   []*Message{},
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Id:         uuid.New(),
	}
}

func (c *Conversation) AddMessage(m *Message) {
	c.Messages = append(c.Messages, m)
}

func (c *Conversation) AddMessages(m []*Message) {
	c.Messages = append(c.Messages, m...)
}
