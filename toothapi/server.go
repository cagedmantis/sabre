package toothapi

import (
	"context"
	"fmt"

	"github.com/cagedmantis/sabre/chunk/store"
)

type toothAPI struct {
	store store.Store
}

func NewToothAPI(s store.Store) ToothAPIServer {
	return &toothAPI{
		store: s,
	}
}

func (t *toothAPI) Write(ctx context.Context, wr *WriteRequest) (*WriteResponse, error) {
	// Check if chunk exists
	// write data to store

	cID := store.ChunkID(wr.ChunkHandle)
	c, err := t.store.GetChunk(cID)
	if err != nil {
		return nil, err
	}

	offset := wr.ByteRange.StartByte
	// confirm that byte range coincides with len(byte)

	b := []byte{}
	// TODO fix these number conversions
	bw, err := c.Write(b, int64(offset))
	if err != nil {
		return nil, err
	}

	if len(b) != bw {
		// TODO in the future, write function to ensure all bytes written
		return nil, fmt.Errorf("Not all bytes written %d/%d", bw, len(b))
	}

	// TODO Ok not needed
	return &WriteResponse{}, nil
}

func (t *toothAPI) Read(ctx context.Context, rr *ReadRequest) (*ReadResponse, error) {
	// Check if chunk exists
	// write data to store

	cID := store.ChunkID(rr.ChunkHandle)
	c, err := t.store.GetChunk(cID)
	if err != nil {
		return nil, err
	}

	offset := rr.ByteRange.StartByte
	// confirm that byte range coincides with len(byte)

	l := rr.ByteRange.EndByte - rr.ByteRange.StartByte

	// TODO fix these number conversions
	b, br, err := c.Read(int(l), int64(offset))
	if err != nil {
		return nil, err
	}

	if int(l) != br {
		// TODO in the future, write function to ensure all bytes written
		return nil, fmt.Errorf("Not all bytes read %d/%d", br, l)
	}

	// TODO change Datum to not be repeated.
	return &ReadResponse{
		Datum: b,
	}, nil
}

func (t *toothAPI) ListChunks(ctx context.Context, lcr *ListChunksRequest) (*ListChunksResponse, error) {
	cs := t.store.ListChunks()

	cIDs := make([]uint64, 0, len(cs))
	for _, id := range cs {
		cIDs = append(cIDs, uint64(id))
	}

	return &ListChunksResponse{
		Chunks: cIDs,
	}, nil
}

func (t *toothAPI) CreateChunk(ctx context.Context, ccr *CreateChunkRequest) (*CreateChunkResponse, error) {
	cID := store.ChunkID(ccr.ChunkHandle)

	if err := t.store.CreateChunk(cID); err != nil {
		return nil, err
	}

	return &CreateChunkResponse{}, nil
}
