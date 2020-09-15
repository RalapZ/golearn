package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fib(filename string) {
	a, b := 0, 1
	defer func() {
		r := recover()
		if errstr, ok := r.(*os.PathError); ok {
			fmt.Println(errstr.Op, errstr.Err, errstr.Path)
			return
		}
	}()
	fileR, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0753)
	//fmt.Printf("%T",fileR)
	//if err!= nil{
	//	return
	//}
	defer fileR.Close()
	fileFlush := bufio.NewWriter(fileR)
	func() {
		for i := 0; i < 20; i++ {
			a, b = b, a+b
			fmt.Println(b)
			fileFlush.WriteString(strconv.Itoa(b) + "\n")
			//fileFlush.Flush()
		}
		fileFlush.Flush()
	}()
}

func FileWriterf(filename string) {
	//fileR,err:=os.OpenFile(filename,os.O_EXCL|os.O_WRONLY,0754)
	//defer func(){
	//	r:=recover()
	//	if errstr,ok:=r.(*os.PathError);ok{
	//		fmt.Println(errstr.Op,errstr.Err,errstr.Path)
	//		return
	//	}
	//}()
	//fileR,err:=os.OpenFile(filename,os.O_EXCL|os.O_WRONLY,0753)
	//fmt.Printf("%T",fileR)
	//if err!=nil{
	//	return
	//}
	//defer fileR.Close()
	fib(filename)

}

func main() {
	//k:=fib()
	filename := "fib.txt"
	FileWriterf(filename)
	//k()
}
