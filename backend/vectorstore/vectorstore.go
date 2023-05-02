package vectorstore

type VectorStore interface {
	AddVector([]float32, any) error
	Query([]float32) (Id string)
}
