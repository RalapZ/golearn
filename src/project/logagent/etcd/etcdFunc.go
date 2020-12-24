package etcdops

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"logagent/model"
	"time"
)

var (
	config             clientv3.Config
	EtcdClient         *clientv3.Client
	watchStartRevision int64
	Etcdconfig         *EtcdConfig
)

type EtcdConfig struct {
	Address string        `ini:"Address"`
	Timeout time.Duration `ini:"Timeout"`
}

func InitETCD(etcdconfig model.Etcdc) (err error) {
	config.Endpoints = etcdconfig.Address
	config.DialTimeout = time.Duration(etcdconfig.Timeout) * time.Millisecond
	EtcdClient, err = clientv3.New(config)
	if err != nil {
		fmt.Println(err)
	}
	return

}
