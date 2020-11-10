package main

import (
	"encoding/json"
	"fmt"
)

type StudentInfo1 struct {
	PID       int
	PPID      int
	name      string
	ChildNode []*StudentInfo1
}

func main() {
	//a := &StudentInfo1{}
	//fmt.Printf("a=%p\n", a)
	//b := a
	//fmt.Printf("b=%p\n", b)
	//a = &StudentInfo1{1, 2, "ralap", nil}
	//fmt.Printf("a=%p\n", a)
	//fmt.Printf("b=%p\n", b)
	//fmt.Println(b)
	a := &StudentInfo1{1, 2, "ra", nil}
	b := &StudentInfo1{2, 3, "ra", nil}
	c := []*StudentInfo1{
		{1, 2, "dd", nil},
		{23, 5, "rr", nil},
		{22, 3, "ee", nil},
	}
	a.ChildNode = append(a.ChildNode, b)

	a.ChildNode = append(a.ChildNode, &StudentInfo1{2, 3, "zp", nil})
	for k, v := range c {
		fmt.Println(k, v)
		a.ChildNode = append(a.ChildNode, c[k])
	}
	bye, _ := json.Marshal(a)
	fmt.Println(string(bye))
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", a.ChildNode[0])
	fmt.Printf("%p\n", a.ChildNode[2])
	fmt.Printf("%p\n", c[0])
}
