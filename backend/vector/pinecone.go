package vector

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aldarisbm/ltmllm/config"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/google/uuid"
	"github.com/nekomeowww/go-pinecone"
	"github.com/pinecone-io/go-pinecone/pinecone_grpc"
	"io"
	"log"
	"net/http"
)

type Pinecone struct {
	client *pinecone.Client
	index  *pinecone.Index
	cfg    *config.Config
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
		cfg:    cfg,
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
		Namespace: "llm",
	}
	resp, err := p.index.Upsert(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("upserting: %v", err)
	}
	p.index.Close()

	return resp, nil
}

func (p *Pinecone) Upsert(msg string, embeddings []float32) {
	upsertReq := UpsertReq{
		Vectors: []Vector{
			{
				Id:     uuid.New().String(),
				Values: embeddings,
				Metadata: Metadata{
					Message: msg,
				},
			},
		},
		Namespace: "llm",
	}
	url := fmt.Sprintf("https://%s-%s.svc.%s.pinecone.io/vectors/upsert", p.cfg.PineconeConfig.IndexName, p.cfg.PineconeConfig.ProjectName, p.cfg.PineconeConfig.Environment)
	b, err := json.Marshal(upsertReq)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Api-Key", p.cfg.PineconeConfig.APIKey)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("req: %+v\n", req)
	fmt.Printf("%+v\n", string(body))
}

type UpsertReq struct {
	Vectors   []Vector `json:"vectors"`
	Namespace string   `json:"namespace"`
}

type Vector struct {
	Id       string    `json:"id"`
	Values   []float32 `json:"values"`
	Metadata Metadata  `json:"metadata"`
}

type Metadata struct {
	Message string `json:"message"`
}
