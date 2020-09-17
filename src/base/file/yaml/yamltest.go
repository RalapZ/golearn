package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//test:
//	user:
//		- Tom
// 		- Lily
//		- Skay
//	mqtt:
//		host: localhost:1883
//		username: test
//		password: test
//
//	http: {port: "8080", host: "127.0.0.1"}

//host: localhost:1883
//username: test
//password: test
type MqttStr struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type HttpStr struct {
	Host string `yaml:host`
	Port string `yaml:port`
}
type Test struct {
	User []string `yaml:user`
	Mqtt MqttStr  `yaml:mqtt`
	Http HttpStr  `yaml:http`
}
type Info struct {
	Testinfo Test `yaml:"test"`
}

func main() {
	filename := "src/base/file/yaml/test.yaml"
	file, _ := ioutil.ReadFile(filename)
	var infor Info
	yaml.Unmarshal(file, &infor)
	fmt.Println(infor)
}

//func (kafkaCluster *MqttStr) getConf() *MqttStr {
//	//应该是 绝对地址
//	yamlFile, err := ioutil.ReadFile("src/base/file/yaml/test1.yaml")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	//err = yaml.Unmarshal(yamlFile, kafkaCluster)
//	err = yaml.UnmarshalStrict(yamlFile, kafkaCluster)
//
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
//	return kafkaCluster
//}
