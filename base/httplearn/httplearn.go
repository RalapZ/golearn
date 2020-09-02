package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://www.baid.com")
	defer resp.Body.Close()
	info, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(info))

}
