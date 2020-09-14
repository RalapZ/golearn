package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	//服务器端口号
	serverPort string
	//钉钉告警url
	dingdingUrl string
)

//钉钉告警模板
const template = `{
     "msgtype": "text",
     "text": {
         "content": "%s"
     },
     "at": {
         "isAtAll": false
     }
 }`

type AlarmStruct struct {
	ScopeId      int    `json:"scopeId"`
	Scope        string `json:"scope"`
	Name         string `json:"name"`
	Id0          string `json:"id0"`
	Id1          string `json:"id1"`
	RuleName     string `json:"ruleName"`
	AlarmMessage string `json:"alarmMessage"`
	StartTime    int    `json:"startTime"`
}

func init() {
	flag.StringVar(&serverPort, "p", "9201", "server port")
	flag.StringVar(&dingdingUrl, "u", "", "alarm webhook")
}

func sendMsg(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	//fmt.Println(string(body))
	data := string(body)
	var st1 []AlarmStruct
	err := json.Unmarshal([]byte(data), &st1)
	fmt.Println(err)
	fmt.Println(data)
	if err == nil {
		//log.Println(data)
		for _, v := range st1 {
			//fmt.Println(v.Name)
			alarmname := v.Name
			alarmscope := v.Scope
			alarmruleName := v.RuleName
			alarmmessage := v.AlarmMessage
			//projstr := v.AlarmMessage
			str := "服务名称:" + alarmname + "\n告警类型:" + alarmscope + "\n告警规则:" + alarmruleName + "" + "\n告警信息:" + alarmmessage
			log.Println(str)
			fmt.Println(str)

			//拼接钉钉所需的消息格式
			msg := strings.NewReader(fmt.Sprintf(template, str))
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := http.Client{Transport: tr}
			req, err := http.NewRequest("POST", dingdingUrl, msg)
			req.Header.Add("Content-Type", "application/json")
			if err != nil {
				log.Print(err)
				os.Exit(0)
			}
			res, err := client.Do(req)
			if err != nil {
				log.Print(err)
				os.Exit(0)
			}
			defer res.Body.Close()
			resbody, err := ioutil.ReadAll(res.Body)
			response := string(resbody)
			if err != nil {
				log.Print(err)
				os.Exit(0)
			}
			fmt.Println(response)
		}
	}
}

func main() {

	flag.Parse()
	//log.Println("test")
	log.Println("DINGURL", dingdingUrl)
	if dingdingUrl == "" {
		log.Fatal("dingdingUrl cannot be empty usage -h get help. ")
	}

	http.HandleFunc("/alarm", sendMsg)

	//启动web服务器
	if err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), nil); err != nil {
		log.Fatal("server start fatal.", err)
	}

}
