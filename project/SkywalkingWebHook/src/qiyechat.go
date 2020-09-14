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
	//fmt.Println(qistu.Access_token)
	//fmt.Printf("%T",qistu.Access_token)
	return qistu.Access_token
}

func main() {

	corpid := "ww97af1eab5d2add3c"
	corpsecret := "iq2IyxRcY3oCHHTFg2U2o3UQGHzXIWkKIAgfKQFdhxw"
	Token := TokenGet(corpid, corpsecret)

	message := make(map[string]interface{})
	message["touser"] = "abelzhu|zhurongjia"
	message["msgtype"] = "text"
	message["agentid"] = "1000002"
	message["text"] = map[string]interface{}{
		"content": "qiyeWeChat message",
	}
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
