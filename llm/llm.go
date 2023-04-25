package llm

import (
	"github.com/sashabaranov/go-openai"
)

func msg() {
	req := openai.ChatCompletionRequest{
		Model: b.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: b.systemContext,
			},
		},
		N:           1,
		MaxTokens:   0,
		Temperature: 0.5,
		Stream:      true,
	}
}
