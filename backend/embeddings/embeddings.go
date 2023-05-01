package embeddings

import (
	"context"
	"github.com/aldarisbm/ltmllm/config"
	"github.com/sashabaranov/go-openai"
)

type Embedding struct {
	client *openai.Client
}

func NewEmbeddingClient(cfg *config.Config) Embedding {
	c := openai.NewClient(cfg.OpenAIConfig.APIKey)
	return Embedding{
		client: c,
	}
}

func (e *Embedding) CreateEmbedding(ctx context.Context, s string) ([]float32, error) {
	req := openai.EmbeddingRequest{
		Input: []string{s},
		Model: openai.AdaEmbeddingV2,
	}
	resp, err := e.client.CreateEmbeddings(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Data[0].Embedding, nil
}

func (e *Embedding) ListEngines(ctx context.Context) ([]openai.Engine, error) {
	resp, err := e.client.ListEngines(ctx)
	if err != nil {
		return nil, err
	}
	return resp.Engines, nil
}
