package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	C := &http.Client{}
	resq, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	resq.Header.Set("name", "Ralap")
	res, _ := C.Do(resq)
	str, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(str))
}
