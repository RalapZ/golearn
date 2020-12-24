package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	// 初始化一个config
	config := sarama.NewConfig()
	// ack就是确保会不会丢失信息，落到磁盘。
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机生产者的分区，集群机器使用
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	// 创建一个生产者实例，参数写入生产者机器信息，端口为9092，传入配置
	client, err := sarama.NewSyncProducer([]string{"10.10.8.150:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	// 创建生产者发送的消息实例，topic和发送内容
	msg := &sarama.ProducerMessage{}
	msg.Topic = "ralap"
	msg.Value = sarama.StringEncoder("日志测试")
	// 发送消息到kafka，返回分区数标志
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send message failed", err)
		return
	}
	fmt.Printf("pid:%v offset:%v\n", pid, offset)
}
