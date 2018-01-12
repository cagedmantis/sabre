package file

import (
	"fmt"

	"github.com/cagedmantis/sabre/file/kv"
)

const (
	// 4kb
	defaultSegmant = 4096

	// 400 kb
	defaultRange = 409600

	keyFileSystem = "filesystem"
)

type FileSystem interface {
}

type fileSystem struct {
	kv            kv.KV
	Path          string //Can be a path to a file in etcd
	AllocatedSize uint64 // in bytes
	UsedSize      uint64
	replicas      Replica
	files         map[string]File
	lastChunk     int64
}

func NewFileSystem(kv kv.KV) (FileSystem, error) {
	fs := &fileSystem{}

	return fs, nil
	//TODO persistence

	// fsb, err := kv.Get(keyFileSystem)
	// if err != nil {
	// 	return nil, err
	// }

	// var buf bytes.Buffer        // Stand-in for a network connection
	// enc := gob.NewEncoder(&buf) // Will write to network.
	// dec := gob.NewDecoder(&buf) // Will read from network.

	// // // Encode (send) the value.
	// // err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	// // if err != nil {
	// // 	log.Fatal("encode error:", err)
	// // }

	// n, err := buf.Write(fsb)
	// if err != nil {
	// 	return

	// // Decode (receive) the value.
	// var q Q
	// err = dec.Decode(&q)
	// if err != nil {
	// 	log.Fatal("decode error:", err)
	// }
	// fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)

	// // HERE ARE YOUR BYTES!!!!
	// fmt.Println(network.Bytes())

}

// func loadFileSystem() fileSystem {
// }

func (fs *fileSystem) OpenFile(path string, create bool) error {
	// does the file exist?
	f, ok := fs.files[path]
	if !ok && !create {
		return fmt.Errorf("file does not exist")
	}
	if !ok && create {
		f := NewFile(path)
	}
	fs.files[path] = f

	// return ranges - chunks - replicas

	return nil
}

type File struct {
	Path   string
	Size   uint64
	ranges *RangeSystem // TODO Not the optimal solution
}

func NewFile(path string) *File {
	return &File{
		Path:   path,
		Size:   0,
		ranges: NewRangeSystem(),
	}
}
