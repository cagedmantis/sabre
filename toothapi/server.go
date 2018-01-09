package toothapi

import "context"

type toothAPI struct {
}

func NewToothAPI() ToothAPIServer {
	return &toothAPI{}
}

func (t *toothAPI) Write(ctx context.Context, wr *WriteRequest) (*WriteResponse, error) {
	return nil, nil
}

func (t *toothAPI) Read(ctx context.Context, rr *ReadRequest) (*ReadResponse, error) {
	return nil, nil
}

func (t *toothAPI) ListChunks(ctx context.Context, lcr *ListChunksRequest) (*ListChunksResponse, error) {
	return nil, nil
}

func (t *toothAPI) CreateChunk(ctx context.Context, ccr *CreateChunkRequest) (*CreateChunkResponse, error) {
	return nil, nil
}
