package toothapi

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type logToothAPI struct {
	next ToothAPIServer
}

func NewLogToothAPI(next ToothAPIServer) ToothAPIServer {
	return &logToothAPI{
		next: next,
	}
}

func (t *logToothAPI) Write(ctx context.Context, wr *WriteRequest) (*WriteResponse, error) {
	log.WithFields(log.Fields{
		"placeholder": "write",
	}).Info("write request received")

	return t.next.Write(ctx, wr)
}
func (t *logToothAPI) Read(ctx context.Context, rr *ReadRequest) (*ReadResponse, error) {
	log.WithFields(log.Fields{
		"placeholder": "write",
	}).Info("read received")

	return t.next.Read(ctx, rr)
}
func (t *logToothAPI) ListChunks(ctx context.Context, lcr *ListChunksRequest) (*ListChunksResponse, error) {
	log.WithFields(log.Fields{
		"placeholder": "write",
	}).Info("list chunks received")

	return t.ListChunks(ctx, lcr)
}
func (t *logToothAPI) CreateChunk(ctx context.Context, ccr *CreateChunkRequest) (*CreateChunkResponse, error) {
	log.WithFields(log.Fields{
		"placeholder": "write",
	}).Info("create chunk received")

	return t.next.CreateChunk(ctx, ccr)
}
