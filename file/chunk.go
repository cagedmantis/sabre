package file

type Chunk struct {
	id   int64
	size int64
}

func NewChunk(id int64, size int64) *Chunk {
	return &Chunk{id, size}
}
