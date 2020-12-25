package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hpcloud/tail"
	etcdops "logagent/etcd"
	logagent_kafka "logagent/kafka"
	"logagent/logtail"
	"logagent/model"
)

var (
	kafkaconfig   *sarama.Config
	logtailconfig = tail.Config{
		ReOpen:    true,
		Location:  &tail.SeekInfo{0, 2},
		MustExist: false,
		Follow:    true,
		Poll:      true,
	}
	LogFile = "logtail/test.log"
)

func main() {
	//file, err := ini.Load("conf/config.ini")
	//if err != nil {
	//	fmt.Println(err)
	//}
	////etcdops.Etcdconfig=new(etcdops.EtcdConfig)
	//config := new(model.IniConfig)
	//err = file.MapTo(config)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(config.Etcdc.Address)
	////model.Fileconfig

	//etcd
	//etcdops.InitETCD(model.Fileconfig.Etcdc)
	model.InitModel("conf/config.ini")
	etcdops.InitETCD(model.Fileconfig.Etcdc)
	etcdops.WatchConfig(model.Fileconfig.TopicName)
	kafkaconfig = sarama.NewConfig()
	// ack就是确保会不会丢失信息，落到磁盘。
	kafkaconfig.Producer.RequiredAcks = sarama.WaitForAll
	// 随机生产者的分区，集群机器使用
	kafkaconfig.Producer.Partitioner = sarama.NewRandomPartitioner
	kafkaconfig.Producer.Return.Successes = true
	err := logagent_kafka.InitKafka(kafkaconfig)
	if err != nil {
		fmt.Println(err)
	}
	err = logtail.InitTail(LogFile, logtailconfig)
	if err != nil {
		fmt.Println(err)
	}
	go logagent_kafka.SendMessage(logagent_kafka.KafkaClient, "ralap", logtail.Lineschan)
	go logtail.TailLine()
	for {
	}
}
