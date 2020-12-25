package model

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"gopkg.in/ini.v1"
	"log"
)

var (
	Fileconfig              = new(IniConfig)
	ModelConfChan           chan string
	LogtailConfigChangeChan chan LogtailConf
	KafkaConfigChangeChan   chan LogKafkaConf

	EtcdClient         *clientv3.Client
	watchStartRevision int64
	Etcdconfig         *EtcdConfig
)

func InitModel(filename string) (err error) {
	file, err := ini.Load(filename)
	if err != nil {
		fmt.Println(err)
	}
	err = file.MapTo(Fileconfig)
	if err != nil {
		return err
	}
	log.Println(Fileconfig.Etcdc)
	return
}

func KafkaLogtailConfigChange() {
	for {
		config := new(interface{})
		select {
		case config = <-LogtailConfigChangeChan:
			fmt.Println(config)
		case config = <-KafkaConfigChangeChan:
			fmt.Println(config)
		default:
		}
	}
}
