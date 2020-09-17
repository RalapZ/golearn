package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

//tencent:
//	auth:
//		corpid:
//		corpsecret:
//	agentid:test
//	user:
//		- zhurongjia

type AuthStr struct {
	CorpInfo   string `yaml:"corpid"`
	CorpSecret string `yaml:"corpsecret"`
}
type TencentConfig struct {
	Auth    AuthStr  `yaml:"auth"`
	Agentid string   `yaml:"agentid"`
	User    []string `yaml:"user"`
}

type Config struct {
	Tencent TencentConfig `yaml:"tencent"`
}

func (conf *Config) ReadConfig(filename string) *Config {
	//filename:="src/project/SkywalkingWebHook/conf/application.yaml1"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if error, ok := err.(*os.PathError); ok {
			log.Println(error.Op, "file not exist", error.Err)
		}
		return nil
	}
	//var conf *Config
	yaml.Unmarshal(file, &conf)
	return conf
	//var
}

func main() {
	var conf *Config
	filename := "src/project/SkywalkingWebHook/conf/application.yaml"
	conf = conf.ReadConfig(filename)
	fmt.Println(conf)
}
