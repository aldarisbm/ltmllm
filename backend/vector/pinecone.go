package vector

import (
	"context"
	"fmt"
	"github.com/aldarisbm/ltmllm/config"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/google/uuid"
	"github.com/nekomeowww/go-pinecone"
	"github.com/pinecone-io/go-pinecone/pinecone_grpc"
	"log"
)

type Pinecone struct {
	client *pinecone.Client
	index  *pinecone.Index
}

func NewPineconeClient(cfg *config.Config) (Pinecone, error) {
	client, err := pinecone.New(
		pinecone.WithAPIKey(cfg.PineconeConfig.APIKey),
		pinecone.WithEnvironment(cfg.PineconeConfig.Environment),
		pinecone.WithProjectName(cfg.PineconeConfig.ProjectName),
	)
	if err != nil {
		log.Fatal(err)
	}
	index, err := client.Index(context.Background(), cfg.PineconeConfig.IndexName)
	if err != nil {
		return Pinecone{}, fmt.Errorf("getting index: %v", err)
	}
	return Pinecone{
		client: client,
		index:  index,
	}, nil
}
func (p *Pinecone) ListIndexes() {
	indexes, err := p.client.ListIndexes()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Indexes: %+v\n", indexes)
}

func (p *Pinecone) UpsertEmbedding(ctx context.Context, msg string, embeddings []float32) (*pinecone_grpc.UpsertResponse, error) {
	req := pinecone_grpc.UpsertRequest{
		Vectors: []*pinecone_grpc.Vector{
			{
				Id:     uuid.New().String(),
				Values: embeddings,
				Metadata: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"message": {
							Kind: &structpb.Value_StringValue{
								StringValue: msg,
							},
						},
					},
				},
			},
		},
		Namespace: "test",
	}
	resp, err := p.index.Upsert(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("upserting: %v", err)
	}

	return resp, nil
}
