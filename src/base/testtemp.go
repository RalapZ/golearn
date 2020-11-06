package main

import "fmt"

func main() {
	s := "A1" // 分配存储"A1"的内存空间，s结构体里的str指针指向这快内存
	fmt.Printf("%p\n", &s)
	s = "A2" // 重新给"A2"的分配内存空间，s结构体里的str指针指向这快内存
	fmt.Printf("%p\n", &s)
	k := []byte{1} // 分配存储1数组的内存空间，s结构体的array指针指向这个数组。
	fmt.Printf("%p\n", k)
	k = []byte{2} // 将array的内容改为2
	fmt.Printf("%p\n", k)

	//k:=[]string{'a','b','c'}
	//fmt.Printf("%p\n",k)

}
