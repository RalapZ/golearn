package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	res, _ := http.Get("http://www.zhenai.com/zhenghun")
	defer res.Body.Close()
	str, _ := ioutil.ReadAll(res.Body)
	re := regexp.MustCompile("http://www.zhen([^\\/])")
	s := re.FindAllStringSubmatch(string(str), -1)
	fmt.Println(s)

	//fmt.Println(string(str))

}
