package logagent_kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var (
	Config      *sarama.Config
	kafkanode   = []string{"10.10.8.150:9092"}
	KafkaClient sarama.SyncProducer
	err         error
	messagedone = make(chan bool)
)

func InitKafka(config *sarama.Config) error {
	fmt.Println(config)
	Config = sarama.NewConfig()
	// ack就是确保会不会丢失信息，落到磁盘。
	Config.Producer.RequiredAcks = config.Producer.RequiredAcks
	// 随机生产者的分区，集群机器使用
	Config.Producer.Partitioner = config.Producer.Partitioner
	Config.Producer.Return.Successes = config.Producer.Return.Successes
	// 创建一个生产者实例，参数写入生产者机器信息，端口为9092，传入配置
	KafkaClient, err = sarama.NewSyncProducer(kafkanode, Config)
	return err
}

func SendMessage(client sarama.SyncProducer, topic string, messagechan <-chan string) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	for true {
		select {
		case message := <-messagechan:
			msg.Value = sarama.StringEncoder(message)
			partition, offset, e := KafkaClient.SendMessage(msg)
			if e != nil {
				fmt.Println(e)
			}
			fmt.Println(partition, offset)
		case <-messagedone:
			return
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}

//func main() {
//	config:=&sarama.Config{}
//	InitKafka(config)
//	messagech:=make(chan string)
//	go SendMessage(KafkaClient,"ralap",messagech)
//	for i:=0;i<10;i++{
//		messagech<- "test"
//	}
//	messagedone<-true;
//}
