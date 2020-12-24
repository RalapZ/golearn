package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	c := clientv3.Config{
		Endpoints:         []string{"10.10.8.150:2379"},
		DialKeepAliveTime: 1000 * time.Microsecond,
	}
	clie, err := clientv3.New(c)
	if err != nil {
		fmt.Println(err)
	}
	kv := clientv3.NewKV(clie)
	kv.Put(context.Background(), "test", "myzone")

	watcher := clientv3.NewWatcher(clie)
	watcher.Watch()
}
