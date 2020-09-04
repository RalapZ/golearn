package main

import (
	"fmt"
)

func HandlerRecover(){
	defer func(){
		r:=recover()
		//fmt.Println(r(error))
		if r == nil{
			fmt.Println("no error occurred ")
			return
		}
		if err,ok:=r.(error);ok{
			fmt.Println("error occurred ",err)
		}else {
			panic(fmt.Sprintf("I do not know what to do %s",r))
		}
	}()
	//panic(errors.New("this is a test error"))
	fmt.Println("te")
}

func main() {
	HandlerRecover()
}
