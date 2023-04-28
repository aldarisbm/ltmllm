package backend

import (
	"github.com/aldarisbm/ltmllm/backend/llm"
	"github.com/aldarisbm/ltmllm/config"
)

func GetNewBot(cfg *config.Config) llm.ChatBot {
	return llm.NewChatBot(cfg)
}
