package store

import (
	"fmt"
	"os"
)

// chunkSizeBytes is the default size of a chunk file 64mb
const defaultSizeBytes = 67108864

type ChunkID uint64
type ChunkPath string

type Chunk interface {
	ID() ChunkID
	Path() ChunkPath
	Write(b []byte, offset int64) (int, error)
	Read(len int, offset int64) ([]byte, int, error)
	Delete() error
	Close()
}

type chunk struct {
	id    ChunkID
	path  ChunkPath
	file  *os.File
	isNew bool
}

// TODO Create a pool of file handles which can be used.
// where are the handles being closed?

func NewChunk(id ChunkID, dir string) (Chunk, error) {
	path := fmt.Sprintf("%s/%d", dir, id)

	f, new, err := createOrOpenFile(path)
	if err != nil {
		return nil, err
	}

	return &chunk{
		id:    id,
		path:  ChunkPath(path),
		file:  f,
		isNew: new, //TODO when is this needed?
	}, nil
}

func (c *chunk) ID() ChunkID {
	return c.id
}

func (c *chunk) Path() ChunkPath {
	return c.path
}

// TODO Should strea writes
func (c *chunk) Write(b []byte, offset int64) (int, error) {
	return c.file.WriteAt(b, offset)
}

// TODO Should strea reads
func (c *chunk) Read(len int, offset int64) ([]byte, int, error) {
	// Could potentially be a large read, allocating a slice for it will impact performance.
	b := make([]byte, len)
	read, err := c.file.ReadAt(b, offset)
	return b, read, err
}

func (c *chunk) Delete() error {
	return nil
}

func (c *chunk) Close() {
	c.file.Close()
}

func createOrOpenFile(path string) (*os.File, bool, error) {
	var isChunkNew bool

	if _, err := os.Stat(path); os.IsNotExist(err) {
		isChunkNew = true

		f, err := createFile(path)
		return f, isChunkNew, err
	}

	isChunkNew = false
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, isChunkNew, err
	}

	return f, isChunkNew, nil
}

func createFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	// TODO populate file with zeros
	// No way truncate will work. lulz
	err = f.Truncate(defaultSizeBytes)
	return f, err
}
