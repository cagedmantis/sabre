package etcd

import (
	"github.com/cagedmantis/sabre/file/kv"

	etcd "github.com/coreos/etcd/client"
)

type etcdKV struct {
	client etcd.Client
}

func NewKVEtcd(c etcd.Client) kv.KV {
	return &etcdKV{
		client: c,
	}
}
