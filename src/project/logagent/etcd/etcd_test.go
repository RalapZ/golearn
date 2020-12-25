package etcdops

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"logagent/model"
	"testing"
	"time"
)

func TestInitETCD(t *testing.T) {
	config := model.Etcdc{}
	config.Address = []string{"10.10.8.150:2379"}
	config.Timeout = 30
	err := InitETCD(config)
	if err != nil {
		fmt.Println(err)
	}
	//var watcher            clientv3.Watcher
	fmt.Println(*EtcdClient)
	watcher := clientv3.NewWatcher(EtcdClient)

	ctx, cancal := context.WithCancel(context.TODO())
	time.AfterFunc(30*time.Second, func() {
		cancal()
	})

	//watcher:=clientv3.NewWatcher(EtcdClient)
	watchchan := watcher.Watch(ctx, "ralap", clientv3.WithRev(0))
	for w := range watchchan {
		for k, v := range w.Events {
			fmt.Println(k, v)
		}
	}

}
