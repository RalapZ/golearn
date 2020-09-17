package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CorpInfo struct {
	corpid     string
	corpsecret string
}
type QiYchatStu struct {
	Errcod       int    `json:"errocd"`
	Errmsg       string `json:errmsg`
	Access_token string `json:access_token`
	expires_in   int
}

type SkywalkInfo struct {
	ScopeId      int    `json:scopeId`
	Scope        string `json:scope`
	Name         string `json:name`
	id0          string
	id1          string
	RuleName     string `json:ruleName`
	AlarmMessage string `json:alarmMessage`
	StartTime    string `json:startTime`
}

func TokenGet(CorpId string, CorpSecret string) string {
	client := &http.Client{}
	data := make(map[string]interface{})
	data["corpid"] = CorpId
	data["corpsecret"] = CorpSecret
	HeaderStr, err := json.Marshal(data)
	if err != nil {
		return "nil"
	}
	req, _ := http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/gettoken", bytes.NewReader(HeaderStr))
	defer req.Body.Close()
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	qistu := QiYchatStu{}
	_ = json.Unmarshal(body, &qistu)
	return qistu.Access_token
}

func SendMessage(CorpId string, CorpSecret string, M SkywalkInfo) {
	Token := TokenGet(CorpId, CorpSecret)
	client := &http.Client{}
	message := make(map[string]interface{})
	message["touser"] = "abelzhu|zhurongjia"
	message["msgtype"] = "text"
	message["agentid"] = "1000002"
	str := "服务名称:" + M.Name + "\n告警类型:" + M.Scope + "\n告警规则:" + M.RuleName + "" + "\n告警信息:" + M.AlarmMessage
	message["text"] = map[string]interface{}{
		"content": str,
	}
	log.Println(str)
	message["safe"] = "0"
	Mdata, _ := json.Marshal(message)
	//fmt.Println(string(Mdata))
	urlR := "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + Token
	req1, _ := http.NewRequest("POST", urlR, bytes.NewReader(Mdata))
	//req1.Header.("access_token",qistu.Access_token)
	fmt.Println(req1.Header)
	resp1, _ := client.Do(req1)
	MMessage, _ := ioutil.ReadAll(resp1.Body)
	fmt.Println("=============")
	fmt.Println(string(MMessage))
}
func AlarmFunc(w http.ResponseWriter, r *http.Request) {

	test, _ := ioutil.ReadAll(r.Body)
	var DataInfo []SkywalkInfo
	json.Unmarshal(test, &DataInfo)
	//fmt.Println(DataInfo)
	for _, Message := range DataInfo {
		corpid := "ww97af1eab5d2add3c"
		corpsecret := "iq2IyxRcY3oCHHTFg2U2o3UQGHzXIWkKIAgfKQFdhxw"
		SendMessage(corpid, corpsecret, Message)
	}
}

func main() {
	http.HandleFunc("/alarm", AlarmFunc)
	http.ListenAndServe(":9000", nil)
}
