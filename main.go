package main

import (
	"github.com/aldarisbm/ltmllm/backend"
	"github.com/aldarisbm/ltmllm/config"
	"github.com/aldarisbm/ltmllm/frontend"
)

func main() {
	cfg := config.NewConfig()
	// Get a new ChatBot
	cb := backend.GetNewBot(&cfg)
	// Run the frontend
	mw := frontend.NewWindow(&cb)
	mw.ShowAndRun()
}
