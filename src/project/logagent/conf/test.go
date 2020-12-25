package main

import (
	"fmt"
	"gopkg.in/ini.v1"
	etcdops "logagent/etcd"
	"time"
)

type Config struct {
	etcdops.EtcdConfig
}

type Etc1dConfig struct {
	Address   string `ini:"Address"`
	Timeout   int    `ini:"Timeout"`
	TopicName string `ini:"TopicName"`
}
type Note struct {
	Content string
	Cities  []string
}

type Person struct {
	Name string
	Age  int `ini:"age"`
	Male bool
	Born time.Time
	Note
	Created time.Time `ini:"-"`
}

func main() {
	//cfg, err:= ini.Load("conf/test.ini")
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//p := new(Person)
	//cfg.MapTo(p)
	//fmt.Println(p)
	//
	//ini.MapTo(p, "conf/ini")
	//// ...
	//fmt.Println(p)

	cfg1, err1 := ini.Load("conf/config.ini")
	if err1 != nil {
		fmt.Println(err1)
	}
	p1 := new(Config)
	cfg1.MapTo(p1)
	fmt.Println(p1)

	//file,err:=ini.Load("conf/config.ini")
	//if err !=nil{
	//	fmt.Println(err)
	//}
	//etcdops.Etcdconfig=new(etcdops.EtcdConfig)
	//config:=new(etcdops.Config)
	//err=file.MapTo(config)
	//if err !=nil{
}
