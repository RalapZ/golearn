package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpGet(httpurl string) {
	params := url.Values{}
	params.Add("name", "ralap")
	Url, _ := url.Parse(httpurl)
	Url.RawQuery = params.Encode()

	str := Url.String()
	fmt.Println(str)
	a, err := http.Get(str)
	if err != nil {
		panic(err)
	}
	defer a.Body.Close()
	fmt.Println(a.StatusCode)
	b, _ := ioutil.ReadAll(a.Body)
	fmt.Println(string(b))

}

func main() {
	HttpGet("http://www.baidu.com")
	//res, _ := http.Get("http://www.zhenai.com/zhenghun")
	//	//defer res.Body.Close()
	//	////res.Header.
	//	//
	//	//str, _ := ioutil.ReadAll(res.Body)
	//	//re := regexp.MustCompile("http://www.zhen([^\\/])")
	//	//s := re.FindAllStringSubmatch(string(str), -1)
	//	//fmt.Println(s)

	//fmt.Println(string(str))

}
