package main

import (
	"context"
	"fmt"
	"github.com/aldarisbm/go-pinecone"
	"github.com/aldarisbm/ltmllm/backend/embeddings"
	"github.com/aldarisbm/ltmllm/config"
	"github.com/google/uuid"
)

func main() {
	cfg := config.NewConfig()
	ec := embeddings.NewEmbeddingClient(&cfg)
	m := "Mi mama me mima yo mimo a mi mama"
	emb, err := ec.CreateEmbedding(context.Background(), m)
	if err != nil {
		panic(err)
	}

	namespace := cfg.PineconeConfig.Namespace
	pc, err := pinecone.NewIndexClient(
		pinecone.WithAPIKey(cfg.PineconeConfig.APIKey),
		pinecone.WithEnvironment(cfg.PineconeConfig.Environment),
		pinecone.WithProjectName(cfg.PineconeConfig.ProjectName),
		pinecone.WithIndexName(cfg.PineconeConfig.IndexName),
	)

	id := uuid.New()
	params := pinecone.UpsertVectorsParams{
		Vectors: []*pinecone.Vector{{
			ID:       id.String(),
			Values:   emb,
			Metadata: map[string]any{"text": m},
		}},
		Namespace: namespace,
	}
	resp, err := pc.UpsertVectors(context.Background(), params)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.UpsertedCount)
}
