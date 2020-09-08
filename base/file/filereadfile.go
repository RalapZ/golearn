package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "./bash/file/readfile.txt"
	defer func() {
		r := recover()
		if r == nil {
			return
		}
		if fileerr, ok := r.(*os.PathError); ok {
			fmt.Println(fileerr.Op, fileerr.Path, fileerr.Err)
		}
	}()
	FileR, _ := os.OpenFile(filename, os.O_RDONLY, 0755)
	//if err !=nil{
	//	panic
	//}
	fileBuf := bufio.NewReader(FileR)
	for {
		str, err := fileBuf.ReadString('\n') // 循环读取一行
		if err != nil {
			fmt.Println("read string failed, err: ", err)
			return
		}
		fmt.Println("read string is %s: ", str)
	}
	//for {}

	//fmt.Println(string(fileBuf)
}
