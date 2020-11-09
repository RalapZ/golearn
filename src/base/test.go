package main

import "fmt"

type StudentInfo1 struct {
	PID       int
	PPID      int
	name      string
	ChildNode []StudentInfo1
}

func main() {
	a := &StudentInfo1{}
	fmt.Printf("a=%p\n", a)
	b := a
	fmt.Printf("b=%p\n", b)
	a = &StudentInfo1{1, 2, "ralap", nil}
	fmt.Printf("a=%p\n", a)
	fmt.Printf("b=%p\n", b)
	fmt.Println(b)
}
