package store

import (
	"errors"
	"io/ioutil"
	"strconv"
)

type Store interface {
	CreateChunk(id ChunkID) error
	ListChunks() []ChunkID
	GetChunk(id ChunkID) (Chunk, error)
	ChunkExists(id ChunkID) bool
	LoadChunks() error
	Close()
}

type chunkStore struct {
	m   map[ChunkID]Chunk
	dir string
}

func NewDiskStore(dirPath string) (Store, error) {
	m := make(map[ChunkID]Chunk)

	cs, err := loadChunks(dirPath)
	if err != nil {
		return nil, err
	}

	for _, c := range cs {
		m[c.ID()] = c
	}

	return &chunkStore{
		m:   m,
		dir: dirPath,
	}, nil
}

func (ds *chunkStore) CreateChunk(id ChunkID) error {
	if _, ok := ds.m[id]; !ok {
		return errors.New("Chunk already exists")
	}

	c, err := NewChunk(id, ds.dir)
	if err != nil {
		return err
	}

	ds.m[id] = c
	return nil
}

func (ds *chunkStore) ListChunks() []ChunkID {
	cs := make([]ChunkID, len(ds.m))
	for idx, chunk := range ds.m {
		cs[idx] = chunk.ID()
	}
	return cs
}

func (ds *chunkStore) GetChunk(id ChunkID) (Chunk, error) {
	v, ok := ds.m[id]
	if !ok {
		return nil, errors.New("Chunk does not exists")
	}
	return v, nil
}

func (ds *chunkStore) ChunkExists(id ChunkID) bool {
	_, ok := ds.m[id]
	return ok

}

func (ds *chunkStore) LoadChunks() error {
	// TODO is this really needed?
	return nil
}

func (ds *chunkStore) Close() {
	for _, chunk := range ds.m {
		chunk.Close()
	}
}

func loadChunks(dirPath string) ([]Chunk, error) {
	cs := []Chunk{}

	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	// TODO leaks handles
	for _, file := range files {
		if !file.IsDir() {
			id, err := strconv.ParseUint(file.Name(), 10, 64)
			if err != nil {
				return nil, err
			}

			cID := ChunkID(id)
			c, err := NewChunk(cID, dirPath)
			if err != nil {
				return nil, err
			}
			cs = append(cs, c)
		}
	}

	return cs, nil
}
