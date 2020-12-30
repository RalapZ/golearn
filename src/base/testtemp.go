package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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

	a := LogtailConf{TopicName: "ralap", PathName: "test1"}
	b := LogtailConf{TopicName: "ralap", PathName: "test"}
	if a == b {
		fmt.Println("match", a, b)
	} else {
		fmt.Println("mismatch")
	}
	k := []string{"test", "test"}
	v := []string{"test", "test", "11"}
	if ok := reflect.DeepEqual(k, v); ok {
		fmt.Println("match", k, v)
	} else {
		fmt.Println("mismatch")
	}

}
