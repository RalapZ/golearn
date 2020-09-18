//##################################################################################################//
//                   			         ┌─┐       ┌─┐ + +                                          //
//                   			      ┌──┘ ┴───────┘ ┴──┐++                                         //
//                   			      │       ───       │++ + + +                                   //
//                   			      ███████───███████ │+                                          //
//                   			      │       ─┴─       │+                                          //
//                   			      └───┐         ┌───┘                                           //
//                   			          │         │   + +                                         //
//                   			          │         └──────────────┐                                //
//                   			          │                        ├─┐                              //
//                   			          │                        ┌─┘                              //
//                   			          └─┐  ┐  ┌───────┬──┐  ┌──┘  + + + +                       //
//                   			            │ ─┤ ─┤       │ ─┤ ─┤                                   //
//                   			            └──┴──┘       └──┴──┘  + + + +                          //
//                   			      神兽出没               永无BUG                                 //
//   Author: Ralap                                                                                  //
//   Date  : 2020/09/17                                                                             //
//##################################################################################################//
package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

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

func (conf *Config) ReadConfig(filename string) {
	//filename:="src/project/SkywalkingWebHook/conf/application.yaml1"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if error, ok := err.(*os.PathError); ok {
			log.Println(error.Op, "file not exist", error.Err)
		}
	}
	yaml.Unmarshal(file, &conf)
}

func main() {
	var conf Config
	str, _ := os.Getwd()
	fmt.Println(string(str))
	filename := "src/project/SkywalkingWebHook/conf/application.yaml"
	conf.ReadConfig(filename)
	fmt.Println(conf)
}
