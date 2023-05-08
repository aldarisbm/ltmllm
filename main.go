package main

import (
	"github.com/aldarisbm/ltmllm/backend/llm"
	"github.com/aldarisbm/ltmllm/config"
)

func main() {
	cfg := config.NewConfig()
	//
	//db, err := bolt.Open(cfg.DatabaseConfig.Path, 0600, nil)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//ec := embeddings.NewEmbeddingClient(&cfg)
	//m := "Mi mama me mima yo mimo a mi mama"
	//emb, err := ec.CreateEmbedding(context.Background(), m)
	//if err != nil {
	//	panic(err)
	//}
	//
	//namespace := cfg.PineconeConfig.Namespace
	//pc, err := pinecone.NewIndexClient(
	//	pinecone.WithAPIKey(cfg.PineconeConfig.APIKey),
	//	pinecone.WithEnvironment(cfg.PineconeConfig.Environment),
	//	pinecone.WithProjectName(cfg.PineconeConfig.ProjectName),
	//	pinecone.WithIndexName(cfg.PineconeConfig.IndexName),
	//)
	//
	//id := uuid.New()
	//params := pinecone.UpsertVectorsParams{
	//	Vectors: []*pinecone.Vector{{
	//		ID:       id.String(),
	//		Values:   emb,
	//		Metadata: map[string]any{"text": m},
	//	}},
	//	Namespace: namespace,
	//}
	//resp, err := pc.UpsertVectors(context.Background(), params)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resp.UpsertedCount)

	chatBot := llm.NewChatBot(&cfg)
	chatBot.ChatStdInput()
}
