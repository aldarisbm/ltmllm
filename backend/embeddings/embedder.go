package embeddings

// Embedder is an interface for embedding clients
type Embedder interface {
	CreateEmbedding(s string) ([]float32, error)
	ListEngines() ([]string, error)
}
