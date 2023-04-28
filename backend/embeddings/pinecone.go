package embeddings

import (
	"github.com/nekomeowww/go-pinecone"
	"log"
)

func (e *Embeddings) NewPineconeClient() {
	client, err := pinecone.New(
		pinecone.WithAPIKey(e.Cfg.PineconeConfig.APIKey),
		pinecone.WithEnvironment(e.Cfg.PineconeConfig.Environment),
		pinecone.WithProjectName(e.Cfg.PineconeConfig.ProjectName),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	resp, err := client.ListIndexes()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(resp)
}
