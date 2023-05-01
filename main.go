package main

import (
	"context"
	"fmt"
	"github.com/aldarisbm/ltmllm/backend/embeddings"
	"github.com/aldarisbm/ltmllm/backend/vector"
	"github.com/aldarisbm/ltmllm/config"
)

//	cfg := config.NewConfig()
//	// Get a new ChatBot
//	cb := backend.GetNewBot(&cfg)
//	// Run the frontend
//	mw := frontend.NewWindow(&cb)
//	mw.ShowAndRun()
//}

func main() {
	cfg := config.NewConfig()
	ec := embeddings.NewEmbeddingClient(&cfg)
	m := "Mi mama me mima yo mimo a mi mama"
	emb, err := ec.CreateEmbedding(context.Background(), m)
	if err != nil {
		panic(err)
	}
	pc, err := vector.NewPineconeClient(&cfg)
	if err != nil {
		panic(err)
	}
	ups, err := pc.UpsertEmbedding(context.Background(), m, emb)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Upserted: %+v\n", ups)

}
