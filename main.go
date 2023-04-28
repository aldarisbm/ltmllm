package main

import (
	"github.com/aldarisbm/ltmllm/backend/llm"
	"github.com/aldarisbm/ltmllm/config"
	"github.com/aldarisbm/ltmllm/frontend"
)

func main() {
	cfg := config.NewConfig()
	cb := llm.NewChatBot(cfg)

	// Run the frontend
	mw := frontend.NewWindow(cb)
	mw.ShowAndRun()
}
