package main

import (
	"github.com/aldarisbm/ltmllm/backend/llm"
	"github.com/aldarisbm/ltmllm/config"
)

func main() {

	//db, err := bolt.Open(cfg.DatabaseConfig.Path, 0600, nil)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()

	cfg := config.NewConfig()

	chatBot := llm.NewChatBot(
		llm.WithApiKey(cfg.OpenAIConfig.APIKey),
		llm.WithSystemContext(cfg.OpenAIConfig.SystemContext),
		llm.WithModel(cfg.OpenAIConfig.Model),
		llm.WithTemperature(cfg.OpenAIConfig.Temperature),
		llm.WithPrompt(cfg.OpenAIConfig.Prompt),
	)
	chatBot.ChatStdInput()
}
