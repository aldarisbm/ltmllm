package main

import (
	"github.com/aldarisbm/ltmllm/backend/chatbot"
	"github.com/aldarisbm/ltmllm/config"
	memory "github.com/aldarisbm/memory/pkg"
	openai "github.com/aldarisbm/memory/pkg/embeddings/openai"
	pc "github.com/aldarisbm/memory/pkg/vectorstore/pinecone"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	_ = godotenv.Load()
	embedder := openai.NewOpenAIEmbedder(
		openai.WithApiKey(os.Getenv("OPENAI_API_KEY")),
	)

	vectorStore := pc.NewStorer(
		pc.WithApiKey(os.Getenv("PINECONE_API_KEY")),
		pc.WithIndexName(os.Getenv("PINECONE_INDEX_NAME")),
		pc.WithProjectName(os.Getenv("PINECONE_PROJECT_NAME")),
		pc.WithEnvironment(os.Getenv("PINECONE_ENVIRONMENT")),
	)

	ltm := memory.NewMemory(
		memory.WithEmbedder(embedder),
		memory.WithVectorStore(vectorStore),
	)

	cfg := config.NewConfig()

	chatBot := chatbot.New(
		chatbot.WithApiKey(cfg.OpenAIConfig.APIKey),
		chatbot.WithSystemContext(cfg.OpenAIConfig.SystemContext),
		chatbot.WithModel(cfg.OpenAIConfig.Model),
		chatbot.WithTemperature(cfg.OpenAIConfig.Temperature),
		chatbot.WithPrompt(cfg.OpenAIConfig.Prompt),
		chatbot.WithMemory(ltm),
	)
	chatBot.StdInput()
}
