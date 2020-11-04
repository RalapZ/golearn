package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func OpenBufioRStr(file string) {
	fi, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	//str,_:=ioutil.ReadAll(fi)
	//fmt.Println(string(str))
	rd := bufio.NewReader(fi)
	//bufio.ScanLines()
	//n:=make([]byte,4096)
	for {
		//rd.ReadLine()
		str, err := rd.ReadString(' ')
		fmt.Println(str)
		if err != nil || err == io.EOF {
			break
		}

	}
}

func main() {
	//fi,_:=os.Create("test111")
	//defer fi.Close()
	////a,_:=fi.Stat()
	//fmt.Println(os.Getwd())
	//fi.Write([]byte("asdfasdfasdfasdfasdfasdfasdfasdfasdfasdf"))
	//fi,_:=os.Open("test111")
	file := "test111"
	OpenBufioRStr(file)
	// 打开文件
}
