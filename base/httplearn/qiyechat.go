package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {
	client := &http.Client{}
	data := make(map[string]interface{})
	//data:=CorpInfo{}
	data["corpid"] = "ww97af1eab5d2add3c"
	data["corpsecret"] = "iq2IyxRcY3oCHHTFg2U2o3UQGHzXIWkKIAgfKQFdhxw"
	//fmt.Println(data)
	bytesData, _ := json.Marshal(data)
	//fmt.Println(string(bytesData))
	req, _ := http.NewRequest("POST", "https://qyapi.weixin.qq.com/cgi-bin/gettoken", bytes.NewReader(bytesData))
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	qistu := QiYchatStu{}
	_ = json.Unmarshal(body, &qistu)
	fmt.Println(qistu.Access_token)
}
