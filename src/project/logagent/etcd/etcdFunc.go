package etcdops

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"logagent/model"
	"time"
)

var (
	config     clientv3.Config
	EtcdClient = model.EtcdClient
)

func InitETCD(etcdconfig model.Etcdc) (err error) {
	config.Endpoints = etcdconfig.Address
	config.DialTimeout = time.Duration(etcdconfig.Timeout) * time.Millisecond
	EtcdClient, err = clientv3.New(config)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func Watchtopic(topicname string) (watcherchan clientv3.WatchChan) {
	watcher := clientv3.NewWatcher(EtcdClient)
	ctx := context.Background()
	watcherchan = watcher.Watch(ctx, topicname, clientv3.WithRev(0))
	return watcherchan
}

func WatchConfig(topicname string) {
	watchConfchan := Watchtopic(topicname)
	for {
		select {
		case conf := <-watchConfchan:
			for _, changeEvent := range conf.Events {
				//fmt.Println(k, string(v.Kv.Value))
				changeConfig := new(model.LogKafkaConf)
				err := json.Unmarshal(changeEvent.Kv.Value, changeConfig)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(changeConfig)
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}

}
