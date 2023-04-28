package embeddings

import (
	"context"
	"github.com/aldarisbm/ltmllm/config"
	"github.com/sashabaranov/go-openai"
)

type Embeddings struct {
	client *openai.Client
	Cfg    *config.Config
}

func NewEmbedding(cfg *config.Config) Embeddings {
	c := openai.NewClient(cfg.OpenAIConfig.APIKey)
	return Embeddings{
		client: c,
		Cfg:    cfg,
	}
}

func (e *Embeddings) CreateEmbeddings(ctx context.Context, s string) *openai.EmbeddingResponse {
	req := openai.EmbeddingRequest{
		Input: []string{s},
		Model: openai.AdaEmbeddingV2,
	}
	resp, _ := e.client.CreateEmbeddings(ctx, req)
	return &resp
}
