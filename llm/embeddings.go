package llm

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

func (b *ChatBot) CreateEmbeddings(ctx context.Context, s string) *openai.EmbeddingResponse {
	req := openai.EmbeddingRequest{
		Input: []string{s},
		Model: openai.AdaEmbeddingV2,
	}
	resp, _ := b.client.CreateEmbeddings(ctx, req)
	return &resp
}
