package file

type Replica struct {
	name   string
	chunks []int64
}

func NewReplica(name string, chunkIDs []int64) *Replica {
	return &Replica{name, chunkIDs}
}
