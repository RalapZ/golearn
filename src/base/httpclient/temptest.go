package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	t, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(t))
}
