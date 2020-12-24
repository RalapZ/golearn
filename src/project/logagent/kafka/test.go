package logagent_kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	kafkanode := []string{"10.10.8.150:9092"}
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	client, err := sarama.NewSyncProducer(kafkanode, config)
	if err != nil {
		fmt.Println(err)
	}
	msg := &sarama.ProducerMessage{}
	msg.Topic = "ralap"
	msg.Value = sarama.StringEncoder("test")
	client.SendMessage(msg)

}
