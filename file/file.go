package file

type FileSystem struct {
	Path          string //Can be a path to a file in etcd
	AllocatedSize uint64 // in bytes
	UsedSize      uint64
	Replicas
}

type File struct {
	Path string
	Size uint64
}

type Chunk struct {
}

type Replica struct {
	Chunks
}
