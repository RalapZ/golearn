package model

import "time"

type IniConfig struct {
	Etcdc
}
type Etcdc struct {
	Address   []string `ini:"Address"`
	Timeout   int      `ini:"Timeout"`
	TopicName string   `ini:"TopicName"`
}

type LogKafkaConf struct {
	Logtail   []LogtailConf `json:"logtail"`
	KafkaConf `json:"kafka"`
}
type LogtailConf struct {
	TopicName string `json:"topicname"`
	LogPath   string `json:"logpath"`
	Timeout   int    `json.:"timeout"`
}

type KafkaConf struct {
	Address string `json:"endpoint"`
	Timeout int    `json:"timeout"`
}

type EtcdConfig struct {
	Address string        `ini:"Address"`
	Timeout time.Duration `ini:"Timeout"`
}
