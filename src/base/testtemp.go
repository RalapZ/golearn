package main

import (
	"encoding/json"
	"fmt"
)

type LogKafkaConf struct {
	LogtailConf []LogtailConf `json:"logtail"`
	KafkaConf   KafkaConf     `json:"kafka"`
}
type LogtailConf struct {
	TopicName string `json:"topicname"`
	PathName  string `json:"pathname"`
}

type KafkaConf struct {
	Address string `json:"address"`
}

func main() {
	var conf LogKafkaConf
	str1 := `{"logtail":[{"topicname":"ralap","pathname":"xxxx"},{"topicname":"myzone","pathname":"xxxx"}],"kafka":{"address":"10.10.8.150:2379"}}`
	err := json.Unmarshal([]byte(str1), &conf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", conf)

}
