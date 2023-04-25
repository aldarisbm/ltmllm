package llm

import (
	"github.com/nekomeowww/go-pinecone"
	"log"
)

func (b *ChatBot) NewPineconeClient() {
	client, err := pinecone.New(
		pinecone.WithAPIKey(b.Cfg.PineconeConfig.APIKey),
		pinecone.WithEnvironment(b.Cfg.PineconeConfig.Environment),
		pinecone.WithProjectName(b.Cfg.PineconeConfig.ProjectName),
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
