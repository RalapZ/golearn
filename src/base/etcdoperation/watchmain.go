package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
		//kv                 clientv3.KV
		watchStartRevision int64
		watcher            clientv3.Watcher
		watchRespChan      <-chan clientv3.WatchResponse
		watchResp          clientv3.WatchResponse
		event              *clientv3.Event
	)

	config = clientv3.Config{
		Endpoints:   []string{"10.10.8.150:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 建立一个客户端
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	// 用于读写etcd的键值对
	//kv = clientv3.NewKV(client)

	// 模拟etcd中kv的变化
	go func() {
		for {
			//kv.Put(context.TODO(), "name", "lesroad")

			//
			//kv.Delete(context.TODO(), "name")

			time.Sleep(1 * time.Second)
		}
	}()

	// 创建一个监听器
	watcher = clientv3.NewWatcher(client)

	// 启动监听 5秒后关闭
	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(60*time.Second, func() {
		fmt.Println("done")
		cancelFunc()
	})
	watchRespChan = watcher.Watch(ctx, "name", clientv3.WithRev(watchStartRevision))

	// 处理kv变化事件
	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			fmt.Println(event.Type)
			fmt.Println("修改为", string(event.Kv.Value))
			//switch event.Type {
			//case mvccpb.PUT:
			//	fmt.Println("修改为", string(event.Kv.Value))
			//case mvccpb.DELETE:
			//	fmt.Println("删除了", string(event.Kv.Key))
			//}
		}
	}
}
