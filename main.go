package main

import (
	"github.com/aldarisbm/ltmllm/config"
	"github.com/aldarisbm/ltmllm/llm"
)

func main() {
	conf := config.NewConfig()
	cb := llm.NewChatBot(conf)
	cb.Chat()
}
