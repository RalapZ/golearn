package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

var (
	config = clientv3.Config{
		Endpoints: []string{"10.10.8.150:2379"},
	}
	etcdcli *clientv3.Client
	err     error
)

func InitETCD() {
	etcdcli, err = clientv3.New(config)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	InitETCD()
	watcher := clientv3.NewWatcher(etcdcli)
	ctx, cacal := context.WithCancel(context.TODO())

	time.AfterFunc(10*time.Second, func() {
		cacal()
	})
	watchan := watcher.Watch(ctx, "ralap")
	for kv := range watchan {
		for k, v := range kv.Events {
			fmt.Println(k, v)
		}
	}
}
