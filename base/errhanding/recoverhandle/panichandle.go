package main

import (
	"fmt"
	"log"
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
	//fmt.Println("te")
}
func test(){
	a:=[3]int{1,2,3}
	b:=4
	fmt.Println(a)
	defer func(){
		r:=recover()
		if err,ok:=r.(error);ok{
			log.Println(err)

		}
	}()
	fmt.Println(a[b])

}

func main() {
	HandlerRecover()
	test()
}
