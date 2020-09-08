package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func FileLineReader(test string) {
	FileR, _ := os.OpenFile(test, os.O_RDONLY, os.ModePerm)
	defer FileR.Close()
	LineRead := bufio.NewReader(FileR)
	for {
		str, _, err := LineRead.ReadLine()
		if err == io.EOF {
			break
		}
		// 如下是某些业务逻辑操作
		// 如下代码打印每次读取的文件行内容
		fmt.Println(string(str))
	}
}

//func FileBlockReader(test string){
//	FileR,_:=os.OpenFile(test,os.O_RDONLY,0755)
//	FileR.Close()
//	const buffersize =64
//	t:=make
//	buffer:=ma	ke([]byte,buffersize)
//	t:=os.Stat
//	for{
//		byteRead,err:=FileR.Read(buffer)
//		content=append(content,buffer[:byteRead]...)
//
//	}
//}
func FileReader(test string) {
	FileR, err := os.OpenFile(test, os.O_RDWR, 0755)
	defer FileR.Close()
	fmt.Println(FileR.Stat())
	FileStat, err := FileR.Stat()
	fmt.Printf("%t", FileStat)
	if err == nil {
		FileRbuffer := bufio.NewReader(FileR)
		filString, _ := ioutil.ReadAll(FileRbuffer)
		fmt.Printf("%s", filString)
		//fmt.Println(FileRbuffer.Size())
	}
}
func FileWriter(test string) {
	FileW, err := os.OpenFile(test, os.O_CREATE|os.O_WRONLY, 0755)
	defer FileW.Close()
	if err == nil {
		FileBuffer := bufio.NewWriter(FileW)
		stringwr := "test\n"
		for i := 0; i < 10; i++ {
			FileBuffer.WriteString(stringwr)
		}
		FileBuffer.Flush()
	}
}
func main() {
	fileWriter := "test.txt"
	FileWriter(fileWriter)
	//FileReader(fileWriter)
	FileLineReader(fileWriter)
}
