package kv

type KV interface {
	Get(key string) ([]byte, error)
	Put(key string, val []byte) error
}
