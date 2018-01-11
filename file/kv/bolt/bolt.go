package bolt

import (
	"github.com/cagedmantis/sabre/file/kv"

	bbolt "github.com/coreos/bbolt"
)

const defaultBucket = "sabre_gum"

type boltKV struct {
	db *bbolt.DB
}

func NewKVbolt(db *bbolt.DB) kv.KV {
	return &boltKV{
		db: db,
	}
}

func (b *boltKV) Get(key string) ([]byte, error) {
	var val []byte

	err := b.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(defaultBucket))
		val = b.Get([]byte(key))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return val, nil
}

func (b *boltKV) Put(key string, val []byte) error {
	return b.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(defaultBucket))
		err := b.Put([]byte(key), val)
		return err
	})
}
